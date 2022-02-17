package display

import "github.com/pjover/sam/internal/domain/ports"

type displayService struct {
	dbService ports.DbService
}

func NewDisplayService(dbService ports.DbService) ports.DisplayService {
	return displayService{
		dbService: dbService,
	}
}

func (d displayService) DisplayCustomer(id int) (string, error) {
	customer, err := d.dbService.FindCustomer(id)
	if err != nil {
		return "", err
	}
	return customer.String(), nil
}

func (d displayService) DisplayInvoice(id string) (string, error) {
	invoice, err := d.dbService.FindInvoice(id)
	if err != nil {
		return "", err
	}
	return invoice.String(), nil
}

func (d displayService) DisplayProduct(id string) (string, error) {
	product, err := d.dbService.FindProduct(id)
	if err != nil {
		return "", err
	}
	return product.String(), nil
}
