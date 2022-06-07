package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateMonthReportCmd struct {
	configService   ports.ConfigService
	generateService ports.GenerateService
}

func NewGenerateMonthReportCmd(configService ports.ConfigService, generateService ports.GenerateService) cli.Cmd {
	return generateMonthReportCmd{
		configService:   configService,
		generateService: generateService,
	}
}

func (e generateMonthReportCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeMes",
		Short:       "Genera l'informe de les factures del mes actual o del indicat",
		Long:        "`Genera l'informe de totes les factures del mes actual o del mes indicat",
		Example:     "   generaInformeMes [mm-aaaa]      Genera l'informe de les factures del mes inidicat",
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfm",
			"generainformemes",
			"genera-informe-mes",
			"generarInformeMes",
			"generarinformemes",
			"generar-informe-mes",
			"gcrem",
			"generateMonthReport",
			"generatemonthreport",
			"generate-month-report",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var yearMonth model.YearMonth
			if len(args) > 0 {
				var err error
				yearMonth, err = model.StringToYearMonth(args[0])
				if err != nil {
					yearMonth = e.configService.GetCurrentYearMonth()
				}
			} else {
				yearMonth = e.configService.GetCurrentYearMonth()
			}
			msg, err := e.generateService.MonthReport(yearMonth)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
