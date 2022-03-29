package create

import (
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

func (c createService) CreateCustomer(customer model.Customer) (string, error) {
	fmt.Printf("Crean el client %s\n", customer.String()) // TODO Remove
	//TODO implement me
	panic("implement me")
}

func (c createService) CreateProduct(product model.Product) (string, error) {
	storedProduct, err := c.dbService.FindProduct(product.Id)
	if err == nil {
		return "", fmt.Errorf("el producte amb codi '%s' ja existeix: %s", product.Id, storedProduct.String())
	}

	err = c.dbService.InsertProduct(product)
	if err != nil {
		return "", fmt.Errorf("error guardant el nou producte: %s", err)
	}

	return fmt.Sprintf("Creat el producte %s", product.String()), nil
}
