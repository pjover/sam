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

	changedSince := c.configService.GetTime("reports.lastCustomersCardsUpdated")
	err := c.configService.SetTime("reports.lastCustomersCardsUpdated", c.osService.Now())
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut actualitzar la configuració: %s", err)
	}

	customers, err := c.dbService.FindChangedCustomers(changedSince)
	if err != nil {
		_ = c.configService.SetTime("reports.lastCustomersCardsUpdated", changedSince)
		return "", fmt.Errorf("no s'ha pogut carregar els consumidors des de la base de dades: %s", err)
	}

	return fmt.Sprintf("Generades %d fitxes de clients", len(customers)), nil
}
