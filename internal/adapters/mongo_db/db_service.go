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
		return model.Customer{}, err
	}

	var result bson.D
	coll := client.Database("hobbit").Collection("customer")
	filter := bson.D{{"_id", code}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Customer{}, fmt.Errorf("no s'ha trobat el ckient amb codi %d", code)
		}
		return model.Customer{}, err
	}

	doc, err := bson.Marshal(result)
	if err != nil {
		return model.Customer{}, err
	}

	var customer dbo.Customer
	err = bson.Unmarshal(doc, &customer)
	if err != nil { // TODO Error: error decoding key children.0.birthDate: cannot decode UTC datetime into a string type
		return model.Customer{}, err
	}

	return dbo.Convert(customer), nil
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
