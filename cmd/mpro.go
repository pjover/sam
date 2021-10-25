package cmd

import (
	"sam/adm"
	"strings"

	"github.com/spf13/cobra"
)

var mproCmd = &cobra.Command{
	Use:         "mpro codiProducte",
	Short:       "Mostra les dades d'un producte",
	Long:        "Mostra les dades d'un producte indicant el seu codi",
	Example:     `   mpro age     Mostra les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateProductCode(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode := strings.ToUpper(args[0])
		return adm.DisplayProduct(productCode)
	},
}

func init() {
	rootCmd.AddCommand(mproCmd)
}
