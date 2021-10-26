package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
	"strings"
)

var eproCmd = &cobra.Command{
	Use:         "epro codiProducte",
	Short:       "Edita les dades d'un producte",
	Long:        "Obri un navegador per a editar les dades d'un producte indicant el seu codi",
	Example:     `   epro age     Edita les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases:     []string{"edita-producte"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateProductCode(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode := strings.ToUpper(args[0])
		return adm.EditProduct(productCode)
	},
}

func init() {
	rootCmd.AddCommand(eproCmd)
}
