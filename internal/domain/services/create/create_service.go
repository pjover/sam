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

	err := product.Validate()
	if err != nil {
		return "", err
	}

	storedProduct, err := c.dbService.FindProduct(product.Id())
	if err == nil {
		return "", fmt.Errorf("el producte amb codi '%s' ja existeix: %s", product.Id(), storedProduct.String())
	}

	err = c.dbService.InsertProduct(product)
	if err != nil {
		return "", fmt.Errorf("error guardant el nou producte: %s", err)
	}

	return fmt.Sprintf("Creat el producte %s", product.String()), nil
}

func (c createService) CreateCustomer(customer model.TransientCustomer) (string, error) {
	err := customer.Validate()
	if err != nil {
		return "", err
	}

	sequence, err := c.getNextCustomerSequence()
	if err != nil {
		return "", err
	}

	newCustomer := c.completeCustomer(customer, sequence)

	err = c.updateDatabase(newCustomer, sequence)
	if err != nil {
		return "", fmt.Errorf("error guardant el nou producte: %s", err)
	}

	return fmt.Sprintf("Creat el client %s\n", newCustomer.String()), nil
}

func (c createService) getNextCustomerSequence() (model.Sequence, error) {
	sequence, err := c.dbService.FindSequence(sequence_type.Customer)
	if err != nil {
		return model.Sequence{}, err
	}

	newSequence := model.NewSequence(sequence_type.Customer, sequence.Counter()+1)
	return newSequence, nil
}

func (c createService) completeCustomer(customer model.TransientCustomer, sequence model.Sequence) model.Customer {
	newCustomerId := sequence.Counter()

	var newChildren []model.Child
	for i, child := range customer.Children {
		newChild := model.NewChild(
			newCustomerId*10+i,
			child.Name,
			child.Surname,
			child.SecondSurname,
			child.TaxID,
			child.BirthDate,
			child.Group,
			child.Note,
			true,
		)
		newChildren = append(newChildren, newChild)
	}

	newCustomer := model.NewCustomer(
		newCustomerId,
		true,
		newChildren,
		customer.Adults,
		customer.InvoiceHolder,
		customer.Note,
		customer.Language,
		c.osService.Now(),
	)
	return newCustomer
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
