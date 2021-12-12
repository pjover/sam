package storage

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/model"

	"github.com/spf13/viper"
)

type CustomerStorage struct {
	getManager tuk.HttpGetManager
}

func NewCustomerStorage() CustomerStorage {
	return CustomerStorage{
		tuk.NewHttpGetManager(),
	}
}

func (c CustomerStorage) GetCustomer(code int) (model.Customer, error) {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), code)
	customer := new(model.Customer)

	err := c.getManager.Type(url, customer)
	if err != nil {
		return model.Customer{}, err
	}
	return *customer, nil
}

func (c CustomerStorage) GetChild(code int) (model.Child, error) {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), code/10)
	customer := new(model.Customer)

	err := c.getManager.Type(url, customer)
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
