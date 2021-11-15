package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/generate"
	"sam/util"
	"time"
)

func newGenerateMonthReportCmd(generateManager generate.GenerateManager) *cobra.Command {
	return &cobra.Command{
		Use:   "generaInformeMes [AAAA-MM]",
		Short: "Genera l'informe de les factures del mes",
		Long:  `Genera l'informe de totes les factures del mes actual`,
		Example: `   generaInformeMes      Genera l'informe de les factures del mes
    - si no s'especifica el mes agafa l'actual`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfm",
			"generainformemes", "genera-informe-mes",
			"generarInformeMes", "generarinformemes", "generar-informe-mes",
			"gcrem",
			"generateMonthReport", "generatemonthreport", "generate-month-report",
		},
		Args: RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var yearMonth time.Time
			var err error
			if len(args) == 0 {
				yearMonth = util.SamTimeManager{}.Now()
			} else {
				yearMonth, err = parseYearMonth(args[0])
				if err != nil {
					return err
				}
			}
			msg, err := generateManager.GenerateMonthReport(yearMonth)
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
	rootCmd.AddCommand(cmd)
}
