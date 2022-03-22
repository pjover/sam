package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
)

type CustomersCardsReports struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewCustomersCardsReports(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) CustomersCardsReports {
	return CustomersCardsReports{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (c CustomersCardsReports) Run() (string, error) {

	return fmt.Sprintf("Generades %d fitxes de clients", 33), nil // TODO fix me
}
