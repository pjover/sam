package mongo_db

import (
	"context"
	"fmt"
	"github.com/pjover/sam/internal/adapters/mongo_db/dbo"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	uri := configService.Get("db.server")
	database := configService.Get("db.name")
	d := dbService{configService, ctx, uri, database}
	d.createIndexes()
	return d
}

func (d dbService) createIndexes() {

	client, err := d.open()
	defer d.close(client)
	if err != nil {
		log.Println("Could not connect to database", err)
		return
	}
	collection := client.Database(d.database).Collection("customer")
	indexes, err := collection.Indexes().List(context.TODO())
	if err != nil {
		log.Println("Could not read customer text indexes:", err)
		return
	}
	if indexes.RemainingBatchLength() > 0 {
		return
	}

	log.Println("Creating MongoDB text indexes ...")
	opt := options.Index()
	opt.SetWeights(bson.M{
		"adults.name":    10,
		"adults.surname": 7,
		"children.name":  1007,
	})
	index := mongo.IndexModel{Keys: bson.M{
		"adults.name":    "text",
		"adults.surname": "text",
		"children.name":  "text",
	}, Options: opt}

	if _, err := collection.Indexes().CreateOne(context.TODO(), index); err != nil {
		log.Println("Could not create customer text index:", err)
	}
}

func (d dbService) FindCustomer(code int) (model.Customer, error) {
	var result dbo.Customer
	if err := d.findOne("customer", code, &result, "el client"); err != nil {
		return model.Customer{}, err
	}
	return dbo.ConvertCustomerToModel(result), nil
}

func (d dbService) FindChild(code int) (model.Child, error) {
	var childCode = code / 10
	customer, err := d.FindCustomer(childCode)
	if err != nil {
		return model.Child{}, err
	}

	var child model.Child
	for _, value := range customer.Children {
		if value.Code == code {
			child = value
			break
		}
	}
	if child == (model.Child{}) {
		return model.Child{}, fmt.Errorf("no s'ha trobat l'infant amb codi %d", code)
	}
	return child, nil
}

func (d dbService) FindInvoice(code string) (model.Invoice, error) {
	var result dbo.Invoice
	if err := d.findOne("invoice", code, &result, "la factura"); err != nil {
		return model.Invoice{}, err
	}
	return dbo.ConvertInvoiceToModel(result), nil
}

func (d dbService) FindProduct(code string) (model.Product, error) {
	var result dbo.Product
	if err := d.findOne("product", code, &result, "el producte"); err != nil {
		return model.Product{}, err
	}
	return dbo.ConvertProductToModel(result), nil
}

func (d dbService) findOne(collection string, code interface{}, result interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)
	err = coll.FindOne(context.TODO(), bson.D{{"_id", code}}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no s'ha trobat %s amb codi %s", name, code)
		}
		return fmt.Errorf("llegint %s amb codi %s des de la base de dades: %s", name, code, err)
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

func (d dbService) FindInvoicesByYearMonth(yearMonth string) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{{"yearMonth", yearMonth}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per any i mes"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindInvoicesByCustomer(customerCode int) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{{"customerId", customerCode}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per client"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoicesToModel(results), nil
}

func (d dbService) FindInvoicesByCustomerAndYearMonth(customerCode int, yearMonth string) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"customerId", customerCode}},
				bson.D{{"year", yearMonth}},
			}},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	if err := d.findMany("invoice", filter, findOptions, &results, "factures per client, any i mes"); err != nil {
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

func (d dbService) SearchCustomers(searchText string) ([]model.Customer, error) {
	var results []dbo.Customer
	filter := bson.M{"$text": bson.M{"$search": searchText}}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	if err := d.findMany("customer", filter, findOptions, &results, "clients actius"); err != nil {
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

func (d dbService) FindAllConsumptions() ([]model.Consumption, error) {
	var results []dbo.Consumption
	filter := bson.D{}
	findOptions := options.Find()
	if err := d.findMany("consumption", filter, findOptions, &results, "tots els consums"); err != nil {
		return nil, err
	}
	return dbo.ConvertConsumptionsToModel(results), nil
}

func (d dbService) FindActiveChildConsumptions(code int) ([]model.Consumption, error) {
	var results []dbo.Consumption
	filter := bson.D{{"childCode", code}, {"invoiceId", ""}}
	findOptions := options.Find()
	if err := d.findMany("consumption", filter, findOptions, &results, "consums per infant"); err != nil {
		return nil, err
	}
	return dbo.ConvertConsumptionsToModel(results), nil
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
	err := d.insertMany("consumption", documents, "consum")
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
		return fmt.Errorf("escrivint %s a la base de dades: %s", name, err)
	}
	return nil
}
