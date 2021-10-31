package storage

import (
	"fmt"
	"sam/comm"
	"sam/model"
)

func GetChild(childCode int) (model.Child, error) {
	url := fmt.Sprintf("http://localhost:8080/customers/%d", childCode/10)
	customer := new(model.Customer)
	err := comm.GetType(url, customer)
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
