package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateMonthInvoices struct {
	generateService ports.GenerateService
}

func NewGenerateMonthInvoices(generateService ports.GenerateService) cli.Cmd {
	return generateMonthInvoices{
		generateService: generateService,
	}
}

func (g generateMonthInvoices) Cmd() *cobra.Command {
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
			msg, err := g.generateService.MonthInvoices()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
