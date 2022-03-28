package create

import (
	"encoding/json"
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
)

type createService struct {
	dbService ports.DbService
}

func NewCreateService(dbService ports.DbService) ports.CreateService {
	return createService{
		dbService: dbService,
	}
}

func (c createService) CreateCustomer(customerJson []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c createService) CreateProduct(productJson []byte) (string, error) {

	var newProduct model.Product
	err := json.Unmarshal(productJson, &newProduct)
	if err != nil {
		return "", fmt.Errorf("error llegint el JSON del nou producte: %s", err)
	}

	storedProduct, err := c.dbService.FindProduct(newProduct.Id)
	if err == nil {
		return "", fmt.Errorf("el producte amb codi '%s' ja existeix: %s", newProduct.Id, storedProduct.String())
	}

	err = c.dbService.InsertProduct(newProduct)
	if err != nil {
		return "", fmt.Errorf("error guardant el nou producte: %s", err)
	}

	return fmt.Sprintf("Creat el producte %s", newProduct.String()), nil
}
