package mongo_db

import (
	"context"
	"fmt"
	"github.com/pjover/sam/internal/adapters/mongo_db/dbo"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
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
	uri := configService.Get("db.server")
	database := configService.Get("db.name")
	return dbService{configService, ctx, uri, database}
}

func (d dbService) FindCustomer(code int) (model.Customer, error) {
	var result dbo.Customer
	if err := d.findOne("customer", code, &result, "el client"); err != nil {
		return model.Customer{}, err
	}
	return dbo.ConvertCustomer(result), nil
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
	return dbo.ConvertInvoice(result), nil
}

func (d dbService) FindProduct(code string) (model.Product, error) {
	var result dbo.Product
	if err := d.findOne("product", code, &result, "el producte"); err != nil {
		return model.Product{}, err
	}
	return dbo.ConvertProduct(result), nil
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
	if err := d.findMany("product", filter, &results, "tots els productes"); err != nil {
		return nil, err
	}
	return dbo.ConvertProducts(results), nil
}

func (d dbService) FindInvoicesByYearMonth(yearMonth string) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"year", yearMonth},
	}
	if err := d.findMany("invoice", filter, &results, "factures per any i mes"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoices(results), nil
}

func (d dbService) FindInvoicesByCustomer(customerCode string) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"customerId", customerCode},
	}
	if err := d.findMany("invoice", filter, &results, "factures per client"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoices(results), nil
}

func (d dbService) FindInvoicesByCustomerAndYearMonth(customerCode string, yearMonth string) ([]model.Invoice, error) {
	var results []dbo.Invoice
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"customerId", customerCode}},
				bson.D{{"year", yearMonth}},
			}},
	}
	if err := d.findMany("invoice", filter, &results, "factures per client, any i mes"); err != nil {
		return nil, err
	}
	return dbo.ConvertInvoices(results), nil
}

func (d dbService) FindActiveCustomers() ([]model.Customer, error) {
	var results []dbo.Customer
	filter := bson.D{
		{"active", true},
	}
	if err := d.findMany("customer", filter, &results, "clients actius"); err != nil {
		return nil, err
	}
	return dbo.ConvertCustomers(results), nil
}

func (d dbService) findMany(collection string, filter bson.D, results interface{}, name string) error {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return fmt.Errorf("connectant a la base de dades: %s", err)
	}

	coll := client.Database(d.database).Collection(collection)
	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("llegint %s des de la base de dades: %s", name, err)
	}

	err = cur.All(context.Background(), results)
	if err != nil {
		return fmt.Errorf("decodificant %s: %s", name, err)
	}
	return nil
}
