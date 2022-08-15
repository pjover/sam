package internal

import (
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/cli"
	adminCli "github.com/pjover/sam/internal/adapters/cli/admin"
	billingCli "github.com/pjover/sam/internal/adapters/cli/billing"
	createCli "github.com/pjover/sam/internal/adapters/cli/create"
	displayCli "github.com/pjover/sam/internal/adapters/cli/display"
	editCli "github.com/pjover/sam/internal/adapters/cli/edit"
	generateCli "github.com/pjover/sam/internal/adapters/cli/generate"
	listCli "github.com/pjover/sam/internal/adapters/cli/list"
	"github.com/pjover/sam/internal/adapters/mongo_express"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/admin"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/pjover/sam/internal/domain/services/create"
	"github.com/pjover/sam/internal/domain/services/display"
	"github.com/pjover/sam/internal/domain/services/edit"
	"github.com/pjover/sam/internal/domain/services/generate"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/pjover/sam/internal/domain/services/list"
	"github.com/pjover/sam/internal/domain/services/loader"
	"github.com/pjover/sam/test/fakes"
)

func InjectDependencies() ports.CommandManager {

	configService := cfg.NewConfigService()
	cmdManager := cli.NewCommandManager(configService)
	dbService := fakes.FakeDbService()
	osService := os.NewOsService()
	langService := lang.NewCatLangService()

	adminServiceDI(cmdManager, configService, osService, langService)
	editServiceDI(cmdManager, configService, osService)
	displayServiceDI(cmdManager, dbService)
	generateServiceDI(cmdManager, configService, dbService, osService, langService)
	listServiceDI(cmdManager, configService, dbService)
	billingServiceDI(cmdManager, configService, dbService, osService)
	createServiceDI(cmdManager, configService, dbService, osService)

	return cmdManager
}

func adminServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, osService ports.OsService, langService lang.LangService) {
	adminService := admin.NewAdminService(configService, osService, langService)
	cmdManager.AddCommand(adminCli.NewBackupCmd(adminService))
}

func editServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, osService ports.OsService) {
	externalEditor := mongo_express.NewExternalEditor(configService, osService)
	editService := edit.NewEditService(externalEditor)
	cmdManager.AddCommand(editCli.NewEditCustomerCmd(editService))
	cmdManager.AddCommand(editCli.NewEditInvoiceCmd(editService))
	cmdManager.AddCommand(editCli.NewEditProductCmd(editService))
}

func displayServiceDI(cmdManager ports.CommandManager, dbService ports.DbService) {
	displayService := display.NewDisplayService(dbService)
	cmdManager.AddCommand(displayCli.NewDisplayCustomerCmd(displayService))
	cmdManager.AddCommand(displayCli.NewDisplayInvoiceCmd(displayService))
	cmdManager.AddCommand(displayCli.NewDisplayProductCmd(displayService))
}

func generateServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService, langService lang.LangService) {
	generateService := generate.NewGenerateService(configService, dbService, osService, langService)
	cmdManager.AddCommand(generateCli.NewGenerateCustomerReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateMonthReportCmd(configService, generateService))
	cmdManager.AddCommand(generateCli.NewGenerateProductReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateSingleInvoice(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateMonthInvoices(configService, generateService))
	cmdManager.AddCommand(generateCli.NewGenerateBddFileCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateCustomerCardsReportsCmd(generateService))
}

func listServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, dbService ports.DbService) {
	bulkLoader := loader.NewBulkLoader(configService, dbService)
	listService := list.NewListService(configService, dbService, bulkLoader)
	cmdManager.AddCommand(listCli.NewListProductsCmd(listService))
	cmdManager.AddCommand(listCli.NewListInvoicesCmd(configService, listService))
	cmdManager.AddCommand(listCli.NewListCustomersCmd(listService))
	cmdManager.AddCommand(listCli.NewListChildrenCmd(listService))
	cmdManager.AddCommand(listCli.NewListMailsCmd(listService))
	cmdManager.AddCommand(listCli.NewListConsumptionsCmd(listService))
}

func billingServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) {
	billingService := billing.NewBillingService(configService, dbService, osService)
	cmdManager.AddCommand(billingCli.NewInsertConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewBillConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewRectifyConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewRectifyConsumptionsCmd(billingService))
}

func createServiceDI(cmdManager ports.CommandManager, configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) {
	createService := create.NewCreateService(dbService, osService)
	cmdManager.AddCommand(createCli.NewCreateProductCmd(createService, configService, osService))
	cmdManager.AddCommand(createCli.NewCreateCustomerCmd(createService, configService, osService))
}
