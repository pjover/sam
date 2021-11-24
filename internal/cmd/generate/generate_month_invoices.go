package generate

import (
	"fmt"

	"github.com/pjover/sam/internal/generate/invoices"
	"github.com/spf13/cobra"
)

var onlyNew bool

func NewGenerateMonthInvoicesCmd() *cobra.Command {
	command := newGenerateMonthInvoicesCmd(invoices.NewMonthInvoicesGenerator())
	command.Flags().BoolVarP(&onlyNew, "nomes_noves", "n", true, "Genera les factures noves, que no s'han generat abans")
	return command
}

func newGenerateMonthInvoicesCmd(generator invoices.MonthInvoicesGenerator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaFactures",
		Short:       "Genera els PDFs de les factures del mes",
		Long:        "Genera els PDFs de les factures del mes al directori 'factures' dedins del directori de treball",
		Example:     `   generaFactures      Genera els PDFs de les factures del mes`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"gfacs",
			"generafactures",
			"genera-factures",
			"generarFactures",
			"generarfactures",
			"generar-factures",
			"ginvs",
			"generateMonthInvoices",
			"generatemonthinvoices",
			"generate-month-invoices",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generator.Generate(onlyNew)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
