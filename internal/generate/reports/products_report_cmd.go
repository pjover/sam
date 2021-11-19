package reports

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/cmd"
)

func newGenerateProductsReportCmd(generator ReportGenerator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeProductes",
		Short:       "Genera l'informe dels productes",
		Long:        "Genera l'informe dels productes actius",
		Example:     `   generaInformeProductes      Genera l'informe dels productes`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfp",
			"generainformeproductes", "genera-informe-productes",
			"generarInformeProductes", "generarinformeproductes", "generar-informe-productes",
			"gcrep",
			"generateProductsReport", "generateproductsreport", "generate-products-report",
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
	generator := NewProductsReportGenerator()
	cmd.RootCmd.AddCommand(newGenerateProductsReportCmd(generator))
}
