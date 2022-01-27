package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateProductReportCmd struct {
	generateService ports.GenerateService
}

func NewGenerateProductReportCmd(generateService ports.GenerateService) cli.Cmd {
	return generateProductReportCmd{
		generateService: generateService,
	}
}

func (e generateProductReportCmd) Cmd() *cobra.Command {
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
			msg, err := e.generateService.ProductReport()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
