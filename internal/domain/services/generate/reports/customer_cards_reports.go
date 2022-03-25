package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"time"
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
		return c.revertLastCustomersCardsUpdated(changedSince, fmt.Errorf("no s'ha pogut carregar els consumidors des de la base de dades: %s", err))
	}

	reportsDir, err := c.configService.GetCustomersCardsDirectory()
	if err != nil {
		return "", err
	}

	for _, customer := range customers {
		err = c.run(reportsDir, customer)
		if err != nil {
			return c.revertLastCustomersCardsUpdated(changedSince, err)
		}
	}

	return fmt.Sprintf("Generades %d fitxes de clients", len(customers)), nil
}

func (c CustomerCardsReports) revertLastCustomersCardsUpdated(changedSince time.Time, err error) (string, error) {
	_ = c.configService.SetTime("reports.lastCustomersCardsUpdated", changedSince)
	return "", err
}

func (c CustomerCardsReports) run(reportsDir string, customer model.Customer) error {
	return nil
}
