package storage

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/model"
	"sam/util"
)

func GetChild(childCode int) (model.Child, error) {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), childCode/10)
	customer := new(model.Customer)
	err := util.GetType(url, customer)
	if err != nil {
		return model.Child{}, err
	}

	var child model.Child
	for _, value := range customer.Children {
		if value.Code == childCode {
			child = value
			break
		}
	}
	if child == (model.Child{}) {
		return model.Child{}, fmt.Errorf("No s'ha trobat l'infant amb codi %d", childCode)
	}
	return child, nil
}