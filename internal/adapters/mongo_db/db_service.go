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
}

func NewDbService(configService ports.ConfigService) ports.DbService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := configService.Get("db.server")
	return dbService{configService, ctx, uri}
}

func (d dbService) open() (*mongo.Client, error) {
	return mongo.Connect(d.ctx, options.Client().ApplyURI(d.uri))
}

func (d dbService) close(client *mongo.Client) {
	if err := client.Disconnect(d.ctx); err != nil {
		panic(err)
	}
}

func (d dbService) GetCustomer(code int) (model.Customer, error) {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return model.Customer{}, fmt.Errorf("error connectant a la base de dades: %s", err)
	}

	var result bson.D
	coll := client.Database("hobbit").Collection("customer")
	filter := bson.D{{"_id", code}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Customer{}, fmt.Errorf("no s'ha trobat el client amb codi %d", code)
		}
		return model.Customer{}, fmt.Errorf("llegint el client %d des de la base de dades: %s", code, err)
	}

	doc, err := bson.Marshal(result)
	if err != nil {
		return model.Customer{}, fmt.Errorf("decodificant el client %d: %s", code, err)
	}

	var customer dbo.Customer
	err = bson.Unmarshal(doc, &customer)
	if err != nil {
		return model.Customer{}, fmt.Errorf("deserialitzant el client %d: %s", code, err)
	}

	return dbo.ConvertCustomer(customer), nil
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
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return model.Invoice{}, fmt.Errorf("connectant a la base de dades: %s", err)
	}

	var result bson.D
	coll := client.Database("hobbit").Collection("invoice")
	filter := bson.D{{"_id", code}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Invoice{}, fmt.Errorf("no s'ha trobat la factura amb codi %s", code)
		}
		return model.Invoice{}, fmt.Errorf("llegint la factura %s des de la base de dades: %s", code, err)
	}

	doc, err := bson.Marshal(result)
	if err != nil {
		return model.Invoice{}, fmt.Errorf("decodificant la factura %s: %s", code, err)
	}

	var invoice dbo.Invoice
	err = bson.Unmarshal(doc, &invoice)
	if err != nil {
		return model.Invoice{}, fmt.Errorf("deserialitzant la factura %s: %s", code, err)
	}

	return dbo.ConvertInvoice(invoice), nil
}

func (d dbService) GetProduct(code string) (model.Product, error) {
	client, err := d.open()
	defer d.close(client)
	if err != nil {
		return model.Product{}, fmt.Errorf("connectant a la base de dades: %s", err)
	}

	var result bson.D
	coll := client.Database("hobbit").Collection("product")
	filter := bson.D{{"_id", code}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Product{}, fmt.Errorf("no s'ha trobat el producte amb codi %s", code)
		}
		return model.Product{}, fmt.Errorf("llegint el producte %s des de la base de dades: %s", code, err)
	}

	doc, err := bson.Marshal(result)
	if err != nil {
		return model.Product{}, fmt.Errorf("decodificant el producte %s: %s", code, err)
	}

	var product dbo.Product
	err = bson.Unmarshal(doc, &product)
	if err != nil {
		return model.Product{}, fmt.Errorf("deserialitzant el producte %s: %s", code, err)
	}

	return dbo.ConvertProduct(product), nil
}
