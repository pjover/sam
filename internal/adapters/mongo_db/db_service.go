package mongo_db

import (
	"context"
	"fmt"
	"github.com/pjover/sam/internal/adapters/mongo_db/dbo"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type dbService struct {
	configService ports.ConfigService
	ctx           context.Context
	uri           string
	database      string
}

func NewDbService(configService ports.ConfigService) ports.DbService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := configService.GetString("db.server")
	database := configService.GetString("db.name")
	d := dbService{configService, ctx, uri, database}
	return d
}

func (d dbService) FindCustomer(id int) (model.Customer, error) {
	var result dbo.Customer
	if err := d.findOne("customer", id, &result, "el client"); err != nil {
		return model.Customer{}, err
	}
	return dbo.ConvertCustomerToModel(result), nil
}

func (d dbService) FindChild(id int) (model.Child, error) {
	var childId = id / 10
	customer, err := d.FindCustomer(childId)
	if err != nil {
		return model.Child{}, err
	}

	var child model.Child
	for _, value := range customer.Children {
		if value.Id == id {
			child = value
			break
		}
	}
	if child == (model.Child{}) {
		return model.Child{}, fmt.Errorf("no s'ha trobat l'infant amb codi %d", id)
	}
	return child, nil
}

func (d dbService) FindInvoice(id string) (model.Invoice, error) {
	var result dbo.Invoice
	if err := d.findOne("invoice", id, &result, "la factura"); err != nil {
		return model.Invoice{}, err
	}
	return dbo.ConvertInvoiceToModel(result), nil
}

func (d dbService) FindProduct(id string) (model.Product, error) {
	var result dbo.Product
	if err := d.findOne("product", id, &result, "el producte"); err != nil {
		return model.Product{}, err
	}
	return dbo.ConvertProductToModel(result), nil
}

func (d dbService) findOne(collection string, id interface{}, result interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)
	err = coll.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no s'ha trobat %s amb codi %s", name, id)
		}
		return fmt.Errorf("llegint %s amb codi %s des de la base de dades: %s", name, id, err)
	}
	return nil
}

func (d dbService) open() (*mongo.Client, error) {
	return mongo.Connect(d.ctx, options.Client().ApplyURI(d.uri))
}

func (d dbService) close(client *mongo.Client) {
	if err := client.Disconnect(d.ctx); err != nil {
		panic(err)
	}
}

func (d dbService) FindAllProducts() ([]model.Product, error) {
	var results []dbo.Product
	filter := bson.D{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("product", filter, findOptions, &results, "tots els productes"); err != nil {
		return nil, err
	}
	return dbo.ConvertProductsToModel(results), nil
}

func (d dbService) FindInvoicesByYearMonth(yearMonth model.YearMonth) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{{"yearMonth", yearMonth.String()}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"customerId", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per any i mes"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindInvoicesByCustomer(customerId int) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{{"customerId", customerId}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per client"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth model.YearMonth) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"customerId", customerId}},
				bson.D{{"yearMonth", yearMonth.String()}},
			}},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per client, any i mes"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindInvoicesByYearMonthAndPaymentTypeAndSentToBank(yearMonth model.YearMonth, paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"yearMonth", yearMonth.String()}},
				bson.D{{"paymentType", dbo.PaymentTypes[paymentType]}},
				bson.D{{"sentToBank", sentToBank}},
			}},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per any i mes, tipus de pagament i enviades al bank"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindActiveCustomers() ([]model.Customer, error) {
	var results []dbo.Customer
	filter := bson.D{{"active", true}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("customer", filter, findOptions, &results, "clients actius"); err != nil {
		return nil, err
	}
	return dbo.ConvertCustomersToModel(results), nil
}

func (d dbService) FindChangedCustomers(changedSince time.Time) ([]model.Customer, error) {
	var results []dbo.Customer
	filter := bson.D{{"changedOn", bson.D{{"$gt", changedSince}}}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("customer", filter, findOptions, &results, "tots els clients"); err != nil {
		return nil, err
	}
	return dbo.ConvertCustomersToModel(results), nil
}

func (d dbService) FindActiveChildren() ([]model.Child, error) {
	customers, err := d.FindActiveCustomers()
	if err != nil {
		return nil, err
	}

	var children []model.Child
	for _, customer := range customers {
		for _, child := range customer.Children {
			if child.Active {
				children = append(children, child)
			}
		}
	}
	return children, nil
}

func (d dbService) FindAllActiveConsumptions() ([]model.Consumption, error) {
	var results []dbo.Consumption
	filter := bson.D{{"invoiceId", "NONE"}}
	findOptions := options.Find()
	if err := d.findMany("consumption", filter, findOptions, &results, "tots els consums"); err != nil {
		return nil, err
	}
	return dbo.ConvertConsumptionsToModel(results), nil
}

func (d dbService) FindActiveChildConsumptions(id int) ([]model.Consumption, error) {
	var results []dbo.Consumption
	filter := bson.D{{"childCode", id}, {"invoiceId", "NONE"}}
	findOptions := options.Find()
	if err := d.findMany("consumption", filter, findOptions, &results, "consums per infant"); err != nil {
		return nil, err
	}
	return dbo.ConvertConsumptionsToModel(results), nil
}

func (d dbService) FindAllSequences() ([]model.Sequence, error) {
	var results []dbo.Sequence
	filter := bson.D{}
	findOptions := options.Find()
	if err := d.findMany("sequence", filter, findOptions, &results, "sequències"); err != nil {
		return nil, err
	}
	return dbo.ConvertSequencesToModel(results), nil
}

func (d dbService) findMany(collection string, filter interface{}, findOptions *options.FindOptions, results interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)
	cur, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return fmt.Errorf("llegint %s des de la base de dades: %s", name, err)
	}

	err = cur.All(context.Background(), results)
	if err != nil {
		return fmt.Errorf("decodificant %s: %s", name, err)
	}
	return nil
}

func (d dbService) InsertConsumptions(consumptions []model.Consumption) error {
	documents := dbo.ConvertConsumptionsToDbo(consumptions)
	err := d.insertMany("consumption", documents, "consums")
	if err != nil {
		return err
	}
	return nil
}

func (d dbService) InsertInvoices(invoices []model.Invoice) error {
	documents := dbo.ConvertInvoicesToDbo(invoices)
	err := d.insertMany("invoice", documents, "factures")
	if err != nil {
		return err
	}
	return nil
}

func (d dbService) InsertProduct(product model.Product) error {
	document := dbo.ConvertProductToDbo(product)
	err := d.insertOne('product', document, 'producte')
	if err != nil {
		return err
	}
	return nil
}

func (d dbService) insertMany(collection string, documents []interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)
	_, err = coll.InsertMany(context.TODO(), documents)
	if err != nil {
		return fmt.Errorf("insertant %s a la base de dades: %s", name, err)
	}
	return nil
}

func (d dbService) UpdateSequences(sequences []model.Sequence) error {
	documents := dbo.ConvertSequencesToDbo(sequences)
	return d.updateMany("sequence", documents, "seqüències")
}

func (d dbService) UpdateConsumptions(consumptions []model.Consumption) error {
	documents := dbo.ConvertConsumptionsToDbo(consumptions)
	return d.updateMany("consumption", documents, "consums")
}

func (d dbService) updateMany(collection string, documents []interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)

	for _, document := range documents {
		doc := document.(dbo.Dbo)
		_, err = coll.ReplaceOne(
			context.TODO(),
			bson.M{"_id": doc.GetId()},
			document,
		)

		if err != nil {
			return fmt.Errorf("actualitzant %s a la base de dades: %s", name, err)
		}
	}
	return nil
}
