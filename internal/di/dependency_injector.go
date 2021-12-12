package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/admin"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/edit"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
	"github.com/pjover/sam/internal/core/os"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services"
)

func InjectDependencies(cfgService ports.ConfigService, cmdManager cli.CmdManager) {

	timeManager := os.NewTimeManager()
	fileManager := os.NewFileManager()
	execManager := os.NewExecManager()

	adminService := services.NewAdminService(cfgService, timeManager, fileManager, execManager)
	cmdManager.AddCommand(admin.NewBackupCmd(adminService))
	cmdManager.AddCommand(admin.NewDirectoryCmd(adminService))

	// TODO move to DI
	cmdManager.AddTmpCommand(consum.NewBillConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewInsertConsumptionsCmd())
	cmdManager.AddTmpCommand(consum.NewRectifyConsumptionsCmd())

	cmdManager.AddTmpCommand(display.NewDisplayCustomerCmd())
	cmdManager.AddTmpCommand(display.NewDisplayInvoiceCmd())
	cmdManager.AddTmpCommand(display.NewDisplayProductCmd())

	cmdManager.AddTmpCommand(edit.NewEditCustomerCmd())
	cmdManager.AddTmpCommand(edit.NewEditInvoiceCmd())
	cmdManager.AddTmpCommand(edit.NewEditProductCmd())

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
