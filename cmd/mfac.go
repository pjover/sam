package cmd

import (
	"sam/adm"
	"strings"

	"github.com/spf13/cobra"
)

var mfacCmd = &cobra.Command{
	Use:         "mfac codiFactura",
	Short:       "Mostra les dades d'una fractura",
	Long:        "Mostra les dades d'uan factura indicant el seu codi",
	Example:     `   mfac f-3945     Mostra les dades de la factura F-3945`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	Aliases:     []string{"mostra-factura"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgs(1, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		invoiceCode := strings.ToUpper(args[0])
		return adm.DisplayInvoice(invoiceCode)
	},
}

func init() {
	rootCmd.AddCommand(mfacCmd)
}
