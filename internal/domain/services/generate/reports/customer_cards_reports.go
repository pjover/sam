package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
)

type CustomerCardsReports struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewCustomerCardsReports(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) CustomerCardsReports {
	return CustomerCardsReports{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (c CustomerCardsReports) Run() (string, error) {

	customers, err := c.dbService.FindAllCustomers()
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut carregar els consumidors des de la base de dades: %s", err)
	}
	return fmt.Sprintf("Generades %d fitxes de clients", len(customers)), nil
}
