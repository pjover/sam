package di

import (
	"github.com/pjover/sam/internal/adapters/primary/cli"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/edit"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
)

func MainDependencyInjection() {
	rootCmd := cli.Execute()
	adminServiceDI(rootCmd)

	// TODO move to DI
	rootCmd.AddCommand(consum.NewBillConsumptionsCmd())
	rootCmd.AddCommand(consum.NewInsertConsumptionsCmd())
	rootCmd.AddCommand(consum.NewRectifyConsumptionsCmd())

	rootCmd.AddCommand(display.NewDisplayCustomerCmd())
	rootCmd.AddCommand(display.NewDisplayInvoiceCmd())
	rootCmd.AddCommand(display.NewDisplayProductCmd())

	rootCmd.AddCommand(edit.NewEditCustomerCmd())
	rootCmd.AddCommand(edit.NewEditInvoiceCmd())
	rootCmd.AddCommand(edit.NewEditProductCmd())

	rootCmd.AddCommand(generate.NewGenerateBddCmd())
	rootCmd.AddCommand(generate.NewGenerateCustomersReportCmd())
	rootCmd.AddCommand(generate.NewGenerateMonthInvoicesCmd())
	rootCmd.AddCommand(generate.NewGenerateMonthReportCmd())
	rootCmd.AddCommand(generate.NewGenerateProductsReportCmd())
	rootCmd.AddCommand(generate.NewGenerateSingleInvoiceCmd())

	rootCmd.AddCommand(list.NewListChildrenCmd())
	rootCmd.AddCommand(list.NewListConsumptionsCmd())
	rootCmd.AddCommand(list.NewListCustomersCmd())
	rootCmd.AddCommand(list.NewListInvoicesCmd())
	rootCmd.AddCommand(list.NewListMailsCmd())
	rootCmd.AddCommand(list.NewListProductsCmd())

	rootCmd.AddCommand(search.NewSearchCustomerCmd())
}
