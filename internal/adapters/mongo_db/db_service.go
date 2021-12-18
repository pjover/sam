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

func (d dbService) GetCustomer(code int) (model.Customer, error) {
	var result dbo.Customer
	if err := d.getOne(code, &result, "customer", "el client"); err != nil {
		return model.Customer{}, err
	}
	return dbo.ConvertCustomer(result), nil
}

func (d dbService) GetChild(code int) (model.Child, error) {

	var childCode = code / 10
	customer, err := d.GetCustomer(childCode)
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

func (d dbService) GetInvoice(code string) (model.Invoice, error) {
	var result dbo.Invoice
	if err := d.getOne(code, &result, "invoice", "la factura"); err != nil {
		return model.Invoice{}, err
	}
	return dbo.ConvertInvoice(result), nil
}

func (d dbService) GetProduct(code string) (model.Product, error) {
	var result dbo.Product

	if err := d.getOne(code, &result, "product", "el producte"); err != nil {
		return model.Product{}, err
	}
	return dbo.ConvertProduct(result), nil
}

func (d dbService) getOne(code interface{}, result interface{}, collection string, name string) error {
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
