package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/display"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services"
)

func displayServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	displayService := services.NewDisplayService(dbService)
	cmdManager.AddCommand(display.NewDisplayCustomerCmd(displayService))
	cmdManager.AddCommand(display.NewDisplayInvoiceCmd(displayService))
}
