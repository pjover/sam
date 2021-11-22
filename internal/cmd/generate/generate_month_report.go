package generate

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/generate"
	"github.com/pjover/sam/internal/generate/reports"
	"github.com/spf13/cobra"
)

func newGenerateMonthReportCmd(generator generate.Generator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeMes [AAAA-MM]",
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
			msg, err := generator.Generate()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	generator := reports.NewMonthReportGenerator()
	command := newGenerateMonthReportCmd(generator)
	cmd.RootCmd.AddCommand(command)
}
