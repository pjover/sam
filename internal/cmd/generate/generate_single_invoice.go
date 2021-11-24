package generate

import (
	"fmt"
	"strings"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/generate/invoices"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(NewGenerateSingleInvoiceCmd())
}

func NewGenerateSingleInvoiceCmd() *cobra.Command {
	return newGenerateSingleInvoiceCmd(invoices.NewSingleInvoiceGenerator())
}

func newGenerateSingleInvoiceCmd(generator invoices.SingleInvoiceGenerator) *cobra.Command {
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
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			invoiceCode := strings.ToUpper(args[0])
			msg, err := generator.Generate(invoiceCode)
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
