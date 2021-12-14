package mongo_db

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
)

type dbService struct {
	configService ports.ConfigService
	getManager    tuk.HttpGetManager
}

func NewDbService() ports.DbService {
	return dbService{
		cfg.NewConfigService(),
		tuk.NewHttpGetManager(),
	}
}

func (d dbService) GetCustomer(code int) (model.Customer, error) {
	baseUrl := d.configService.Get("urls.hobbit")
	url := fmt.Sprintf("%s/customers/%d", baseUrl, code)
	customer := new(model.Customer)

	err := d.getManager.Type(url, customer)
	if err != nil {
		return model.Customer{}, err
	}
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
