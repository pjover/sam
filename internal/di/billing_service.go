package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	billingCli "github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
)

func billingServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	billingService := billing.NewBillingService(dbService)
	cmdManager.AddCommand(billingCli.NewInsertConsumptionsCmd(billingService))
}
