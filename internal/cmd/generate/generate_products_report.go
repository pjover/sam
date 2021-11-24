package generate

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/generate"
	"github.com/pjover/sam/internal/generate/reports"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(NewGenerateProductsReportCmd())
}

func NewGenerateProductsReportCmd() *cobra.Command {
	return newGenerateProductsReportCmd(reports.NewProductsReportGenerator())
}

func newGenerateProductsReportCmd(generator generate.Generator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeProductes",
		Short:       "Genera l'informe dels productes",
		Long:        "Genera l'informe dels productes actius",
		Example:     `   generaInformeProductes      Genera l'informe dels productes`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfp",
			"generainformeproductes",
			"genera-informe-productes",
			"generarInformeProductes",
			"generarinformeproductes",
			"generar-informe-productes",
			"gcrep",
			"generateProductsReport",
			"generateproductsreport",
			"generate-products-report",
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
