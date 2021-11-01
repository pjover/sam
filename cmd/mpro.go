package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var mproCmd = &cobra.Command{
	Use:         "mpro codiProducte",
	Short:       "Mostra les dades d'un producte",
	Long:        "Mostra les dades d'un producte indicant el seu codi",
	Example:     `   mpro age     Mostra les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases:     []string{"mostra-producte"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgsEqualsTo(1, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode, err := parseProductCode(args[0])
		if err != nil {
			return err
		}
		return adm.DisplayProduct(productCode)
	},
}

func init() {
	rootCmd.AddCommand(mproCmd)
}
