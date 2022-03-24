package internal

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/admin"
	billingCli "github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/adapters/cli/display"
	"github.com/pjover/sam/internal/adapters/cli/edit"
	generateCli "github.com/pjover/sam/internal/adapters/cli/generate"
	listCli "github.com/pjover/sam/internal/adapters/cli/list"
	searchCmd "github.com/pjover/sam/internal/adapters/cli/search"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/mongo_express"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/domain/ports"
	admin2 "github.com/pjover/sam/internal/domain/services/admin"
	"github.com/pjover/sam/internal/domain/services/billing"
	display2 "github.com/pjover/sam/internal/domain/services/display"
	edit2 "github.com/pjover/sam/internal/domain/services/edit"
	"github.com/pjover/sam/internal/domain/services/generate"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/pjover/sam/internal/domain/services/list"
	"github.com/pjover/sam/internal/domain/services/search"
)

func MainDI(cmdManager cli.CmdManager, configService ports.ConfigService) {

	dbService := mongo_db.NewDbService(configService)
	osService := os.NewOsService()
	langService := lang.NewCatLangService()

	adminServiceDI(cmdManager, configService, osService, langService)
	editServiceDI(cmdManager, configService, osService)
	displayServiceDI(cmdManager, dbService)
	generateServiceDI(cmdManager, configService, dbService, osService, langService)
	listServiceDI(cmdManager, configService, dbService)
	searchServiceDI(cmdManager, dbService)
	billingServiceDI(cmdManager, configService, dbService, osService)
}

func adminServiceDI(cmdManager cli.CmdManager, configService ports.ConfigService, osService ports.OsService, langService lang.LangService) {
	adminService := admin2.NewAdminService(configService, osService, langService)
	cmdManager.AddCommand(admin.NewBackupCmd(adminService))
}

func editServiceDI(cmdManager cli.CmdManager, configService ports.ConfigService, osService ports.OsService) {
	externalEditor := mongo_express.NewExternalEditor(configService, osService)
	editService := edit2.NewEditService(externalEditor)
	cmdManager.AddCommand(edit.NewEditCustomerCmd(editService))
	cmdManager.AddCommand(edit.NewEditInvoiceCmd(editService))
	cmdManager.AddCommand(edit.NewEditProductCmd(editService))
}

func displayServiceDI(cmdManager cli.CmdManager, dbService ports.DbService) {
	displayService := display2.NewDisplayService(dbService)
	cmdManager.AddCommand(display.NewDisplayCustomerCmd(displayService))
	cmdManager.AddCommand(display.NewDisplayInvoiceCmd(displayService))
	cmdManager.AddCommand(display.NewDisplayProductCmd(displayService))
}

func generateServiceDI(cmdManager cli.CmdManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService, langService lang.LangService) {
	generateService := generate.NewGenerateService(configService, dbService, osService, langService)
	cmdManager.AddCommand(generateCli.NewGenerateCustomerReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateMonthReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateProductReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateSingleInvoice(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateMonthInvoices(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateBddFileCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateCustomersCardsReportsCmd(generateService))
}

func listServiceDI(cmdManager cli.CmdManager, configService ports.ConfigService, dbService ports.DbService) {
	listService := list.NewListService(dbService)
	cmdManager.AddCommand(listCli.NewListProductsCmd(listService))
	cmdManager.AddCommand(listCli.NewListInvoicesCmd(configService, listService))
	cmdManager.AddCommand(listCli.NewListCustomersCmd(listService))
	cmdManager.AddCommand(listCli.NewListChildrenCmd(listService))
	cmdManager.AddCommand(listCli.NewListMailsCmd(listService))
	cmdManager.AddCommand(listCli.NewListConsumptionsCmd(listService))
}

func searchServiceDI(cmdManager cli.CmdManager, dbService ports.DbService) {
	searchService := search.NewSearchService(dbService)
	cmdManager.AddCommand(searchCmd.NewSearchCustomerCmd(searchService))
}

func billingServiceDI(cmdManager cli.CmdManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) {
	billingService := billing.NewBillingService(configService, dbService, osService)
	cmdManager.AddCommand(billingCli.NewInsertConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewBillConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewRectifyConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewRectifyConsumptionsCmd(billingService))
}
