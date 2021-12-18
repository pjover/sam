package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/lang"
)

func MainDI(configService ports.ConfigService, cmdManager cli.CmdManager) {

	osService := os.NewOsService()
	langService := lang.NewLangService(configService.Get("lang"))
	dbService := mongo_db.NewDbService(configService)

	adminServiceDI(configService, cmdManager, osService, langService)
	editServiceDI(configService, cmdManager, osService)
	displayServiceDI(dbService, cmdManager)

	// TODO move to DI and remove method AddTmpCommand
	cmdManager.AddTmpCommand(consum.NewBillConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewInsertConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewRectifyConsumptionsCmd())

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
