package mongo_db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type dbService struct {
	configService ports.ConfigService
	getManager    tuk.HttpGetManager
	ctx           context.Context
	uri           string
}

func NewDbService(configService ports.ConfigService) ports.DbService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := configService.Get("db.server")

	return dbService{
		configService,
		tuk.NewHttpGetManager(),
		ctx,
		uri,
	}
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

	coll := client.Database("hobbit").Collection("customer")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"_id", code}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return model.Customer{}, fmt.Errorf("customer %d not found", code)
		}
		return model.Customer{}, err
	}
	// end findOne

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Customer:\n%s\n", output)

	customer := new(model.Customer)
	err = json.Unmarshal([]byte(output), &customer)
	if err != nil {
		return model.Customer{}, err
	}
	//baseUrl := d.configService.Get("urls.hobbit")
	//url := fmt.Sprintf("%s/customers/%d", baseUrl, code)
	//
	//
	//err = d.getManager.Type(url, customer)
	//if err != nil {
	//	return model.Customer{}, err
	//}
	return *customer, nil
}

func (d dbService) GetChild(code int) (model.Child, error) {
	baseUrl := d.configService.Get("urls.hobbit")
	url := fmt.Sprintf("%s/customers/%d", baseUrl, code/10)
	customer := new(model.Customer)

	err := d.getManager.Type(url, customer)
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
		return model.Child{}, fmt.Errorf("No s'ha trobat l'infant amb codi %d", code)
	}
	return child, nil
}
