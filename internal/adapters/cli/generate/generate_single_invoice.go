package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
	"strings"
)

type generateSingleInvoice struct {
	generateService ports.GenerateService
}

func NewGenerateSingleInvoice(generateService ports.GenerateService) cli.Cmd {
	return generateSingleInvoice{
		generateService: generateService,
	}
}

func (g generateSingleInvoice) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "generaFactura",
		Short:       "Genera el PDF d'una factura",
		Long:        "Genera el PDF d'una factura indicant el seu codi",
		Example:     `   generaFactura f-3945     Genera el PDF de la factura F-3945`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"gfac",
			"generafactura",
			"genera-factura",
			"generarFactura",
			"generarfactura",
			"generar-factura",
			"ginv",
			"generateSingleInvoice",
			"generatesingleinvoice",
			"generate-single-invoice",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := strings.ToUpper(args[0])
			msg, err := g.generateService.SingleInvoice(id)
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
