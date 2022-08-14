package e2e

import (
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
)

func addBillingCommands(cmdManager ports.CommandManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) {
	billingService := billing.NewBillingService(configService, dbService, osService)
	cmdManager.AddCommand(NewCommand(billingService, InsertConsumptions, Arguments{"2630", "1", "QME", "2", "MME", "1", "AGE"}))
	cmdManager.AddCommand(NewCommand(billingService, InsertConsumptions, Arguments{"2540", "1", "QME", "1", "MME", "1", "AGE"}))
}
