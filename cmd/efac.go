package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
	"strings"
)

var efacCmd = &cobra.Command{
	Use:         "efac codiFactura",
	Short:       "Edita les dades d'una factura",
	Long:        "Obri un navegador per a editar les dades d'una factura indicant el seu codi",
	Example:     `   efac f-3945     Edita les dades de la factura F-3945`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases:     []string{"edita-factura"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgsEqualsTo(1, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		invoiceCode := strings.ToUpper(args[0])
		return adm.EditInvoice(invoiceCode)
	},
}

func init() {
	rootCmd.AddCommand(efacCmd)
}
