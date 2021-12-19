package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/cobra"
)

type generateMonthReportCmd struct {
	generateService ports.GenerateService
}

func NewGenerateMonthReportCmd(generateService ports.GenerateService) cli.Cmd {
	return generateMonthReportCmd{
		generateService: generateService,
	}
}

func (e generateMonthReportCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeMes",
		Short:       "Genera l'informe de les factures del mes actual",
		Long:        "`Genera l'informe de totes les factures del mes actual",
		Example:     "   generaInformeMes      Genera l'informe de les factures del mes actual",
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
			msg, err := e.generateService.MonthReport()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
