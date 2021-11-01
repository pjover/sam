package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var eproCmd = &cobra.Command{
	Use:         "epro codiProducte",
	Short:       "Edita les dades d'un producte",
	Long:        "Obri un navegador per a editar les dades d'un producte indicant el seu codi",
	Example:     `   epro age     Edita les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases:     []string{"edita-producte"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgsEqualsTo(1, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode, err := parseProductCode(args[0])
		if err != nil {
			return err
		}
		return adm.EditProduct(productCode)
	},
}

func init() {
	rootCmd.AddCommand(eproCmd)
}
