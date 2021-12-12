package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/lang"
)

func MainDI(cfgService ports.ConfigService, cmdManager cli.CmdManager) {

	osService := os.NewOsService()
	langService := lang.NewLangService(cfgService.Get("lang"))

	adminServiceDI(cfgService, cmdManager, osService, langService)
	editServiceDI(cmdManager, osService)

	// TODO move to DI and remove method AddTmpCommand
	cmdManager.AddTmpCommand(consum.NewBillConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewInsertConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewRectifyConsumptionsCmd())

	cmdManager.AddTmpCommand(display.NewDisplayCustomerCmd())
	cmdManager.AddTmpCommand(display.NewDisplayInvoiceCmd())
	cmdManager.AddTmpCommand(display.NewDisplayProductCmd())

	cmdManager.AddTmpCommand(generate.NewGenerateBddCmd())
	cmdManager.AddTmpCommand(generate.NewGenerateCustomersReportCmd())
	cmdManager.AddTmpCommand(generate.NewGenerateMonthInvoicesCmd())
	cmdManager.AddTmpCommand(generate.NewGenerateMonthReportCmd())
	cmdManager.AddTmpCommand(generate.NewGenerateProductsReportCmd())
	cmdManager.AddTmpCommand(generate.NewGenerateSingleInvoiceCmd())

	cmdManager.AddTmpCommand(list.NewListChildrenCmd())
	cmdManager.AddTmpCommand(list.NewListConsumptionsCmd())
	cmdManager.AddTmpCommand(list.NewListCustomersCmd())
	cmdManager.AddTmpCommand(list.NewListInvoicesCmd())
	cmdManager.AddTmpCommand(list.NewListMailsCmd())
	cmdManager.AddTmpCommand(list.NewListProductsCmd())

	cmdManager.AddTmpCommand(search.NewSearchCustomerCmd())
}
