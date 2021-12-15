package display

import (
	"github.com/pjover/sam/internal/core/ports"
	"strconv"
)

type CustomerDisplay struct {
	dbService ports.DbService
}

func NewCustomerDisplay(dbService ports.DbService) Display {
	return CustomerDisplay{
		dbService: dbService,
	}
}

func (c CustomerDisplay) Display(code string) (string, error) {
	codeI, _ := strconv.Atoi(code)
	customer, err := c.dbService.GetCustomer(codeI)
	if err != nil {
		return "", err
	}
	return customer.String(), nil
}
