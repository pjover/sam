package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateMonthInvoices struct {
	configService   ports.ConfigService
	generateService ports.GenerateService
}

func NewGenerateMonthInvoices(configService ports.ConfigService, generateService ports.GenerateService) cli.Cmd {
	return generateMonthInvoices{
		configService:   configService,
		generateService: generateService,
	}
}

func (g generateMonthInvoices) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generaFactures [mm-aaaa]",
		Short: "Genera els PDFs de les factures del mes actual o del indicat",
		Long:  "Genera els PDFs de les factures del mes al directori 'factures' dedins del directori de treball",
		Example: `   generaFactures      Genera els PDFs de les factures del mes actual
   generaFactures  "04-2022"    Genera els PDFs de les factures d'abril de 2022`,
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
			var yearMonth model.YearMonth
			if len(args) > 0 {
				var err error
				yearMonth, err = model.StringToYearMonth(args[0])
				if err != nil {
					yearMonth = g.configService.GetCurrentYearMonth()
				}
			} else {
				yearMonth = g.configService.GetCurrentYearMonth()
			}
			msg, err := g.generateService.MonthInvoices(yearMonth)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
