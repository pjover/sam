package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
	"strings"
)

var editInvoiceCmd = &cobra.Command{
	Use:         "editaFactura codiFactura",
	Short:       "Edita les dades d'una factura",
	Long:        "Obri un navegador per a editar les dades d'una factura indicant el seu codi",
	Example:     `   editaFactura f-3945     Edita les dades de la factura F-3945`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"efac",
		"editafactura", "edita-factura",
		"editarFactura", "editarfactura", "editar-factura",
		"einv",
		"editInvoice", "editinvoice", "edit-invoice",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		invoiceCode := strings.ToUpper(args[0])
		return adm.EditInvoice(invoiceCode)
	},
}

func init() {
	rootCmd.AddCommand(editInvoiceCmd)
}
