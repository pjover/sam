package services

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
	customer, err := d.dbService.GetCustomer(code)
	if err != nil {
		return "", err
	}
	return customer.String(), nil
}
