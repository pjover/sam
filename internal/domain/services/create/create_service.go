package create

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"github.com/pjover/sam/internal/domain/ports"
)

type createService struct {
	dbService ports.DbService
	osService ports.OsService
}

func NewCreateService(dbService ports.DbService, osService ports.OsService) ports.CreateService {
	return createService{
		dbService: dbService,
		osService: osService,
	}
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

func (c createService) CreateCustomer(customer model.Customer) (string, error) {
	err := c.validateCustomer(customer)
	if err != nil {
		return "", err
	}

	sequence, err := c.getNextCustomerSequence()
	if err != nil {
		return "", err
	}

	customer = c.completeCustomer(customer, sequence)

	err = c.updateDatabase(customer, sequence)
	if err != nil {
		return "", fmt.Errorf("error guardant el nou producte: %s", err)
	}

	return fmt.Sprintf("Creat el client %s\n", customer.String()), nil
}

func (c createService) validateCustomer(customer model.Customer) error {
	return nil // TODO Validate mandatory fields, TaxIds & Iban
}

func (c createService) getNextCustomerSequence() (model.Sequence, error) {
	sequence, err := c.dbService.FindSequence(sequence_type.Customer)
	if err != nil {
		return model.Sequence{}, err
	}

	newSequence := model.Sequence{
		Id:      sequence_type.Customer,
		Counter: sequence.Counter + 1,
	}
	return newSequence, nil
}

func (c createService) completeCustomer(customer model.Customer, sequence model.Sequence) model.Customer {
	customer.Id = sequence.Counter

	for i := 0; i < len(customer.Children); i++ {
		customer.Children[i].Id = customer.Id*10 + i
	}

	customer.ChangedOn = c.osService.Now()
	return customer
}

func (c createService) updateDatabase(customer model.Customer, sequence model.Sequence) error {
	err := c.dbService.InsertCustomer(customer)
	if err != nil {
		return err
	}

	err = c.dbService.UpdateSequence(sequence)
	if err != nil {
		return err
	}
	return nil
}
