package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/generate"
	"strings"
)

func newGenerateInvoiceCmd(generateManager generate.GenerateManager) *cobra.Command {
	return &cobra.Command{
		Use:         "generaFactura",
		Short:       "Genera el PDF d'una factura",
		Long:        "Genera el PDF d'una factura indicant el seu codi",
		Example:     `   generaFactura f-3945     Genera el PDF de la factura F-3945`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"gfac",
			"generafactura", "genera-factura",
			"generarFactura", "generarfactura", "generar-factura",
			"ginv",
			"generateInvoice", "generateinvoice", "generate-invoice",
		},
		Args: ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			invoiceCode := strings.ToUpper(args[0])
			msg, err := generateManager.GenerateInvoice(invoiceCode)
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	cmd := newGenerateInvoiceCmd(generate.NewGenerateManager())
	RootCmd.AddCommand(cmd)
}
