package display

import "github.com/pjover/sam/internal/core/ports"

type displayService struct {
	dbService ports.DbService
}

func NewDisplayService(dbService ports.DbService) ports.DisplayService {
	return displayService{
		dbService: dbService,
	}
}

func (d displayService) DisplayCustomer(code int) (string, error) {
	customer, err := d.dbService.FindCustomer(code)
	if err != nil {
		return "", err
	}
	return customer.String(), nil
}

func (d displayService) DisplayInvoice(code string) (string, error) {
	invoice, err := d.dbService.FindInvoice(code)
	if err != nil {
		return "", err
	}
	return invoice.String(), nil
}

func (d displayService) DisplayProduct(code string) (string, error) {
	product, err := d.dbService.FindProduct(code)
	if err != nil {
		return "", err
	}
	return product.String(), nil
}
