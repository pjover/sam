package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/generate"
)

func newGenerateProductsReportCmd(generateManager generate.GenerateManager) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeProductes",
		Short:       "Genera l'informe dels productes",
		Long:        "Genera l'informe dels productes actius",
		Example:     `   generaInformeProductes      Genera l'informe dels productes`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfp",
			"generainfproductes", "genera-inf-productes",
			"generarInfProductes", "generarinfproductes", "generar-inf-productes",
			"gcrep",
			"generateProductsReport", "generateproductsreport", "generate-products-report",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generateManager.GenerateProductsReport()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	cmd := newGenerateProductsReportCmd(generate.NewGenerateManager())
	rootCmd.AddCommand(cmd)
}
