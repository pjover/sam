package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/display"
	"github.com/pjover/sam/internal/core/ports"
	display2 "github.com/pjover/sam/internal/core/services/display"
)

func displayServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	displayService := display2.NewDisplayService(dbService)
	cmdManager.AddCommand(display.NewDisplayCustomerCmd(displayService))
	cmdManager.AddCommand(display.NewDisplayInvoiceCmd(displayService))
	cmdManager.AddCommand(display.NewDisplayProductCmd(displayService))
}
