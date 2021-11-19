package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/generate"
)

func newGenerateMonthReportCmd(generateManager generate.GenerateManager) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeMes [AAAA-MM]",
		Short:       "Genera l'informe de les factures del mes actual",
		Long:        "`Genera l'informe de totes les factures del mes actual",
		Example:     "   generaInformeMes      Genera l'informe de les factures del mes actual",
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfm",
			"generainformemes", "genera-informe-mes",
			"generarInformeMes", "generarinformemes", "generar-informe-mes",
			"gcrem",
			"generateMonthReport", "generatemonthreport", "generate-month-report",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generateManager.GenerateMonthReport()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	cmd := newGenerateMonthReportCmd(generate.NewGenerateManager())
	RootCmd.AddCommand(cmd)
}
