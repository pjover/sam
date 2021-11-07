package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var displayProductCmd = &cobra.Command{
	Use:         "mostraProducte codiProducte",
	Short:       "Mostra les dades d'un producte",
	Long:        "Mostra les dades d'un producte indicant el seu codi",
	Example:     `   mostraProducte age     Mostra les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"mpro",
		"mostraproducte", "mostra-producte",
		"mostrarProducte", "mostrarproducte", "mostrar-producte",
		"dpro",
		"displayProduct", "displayproduct", "display-product",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode, err := parseProductCode(args[0])
		if err != nil {
			return err
		}

		manager := adm.NewDisplayManager()
		_, err = manager.DisplayProduct(productCode)
		return err
	},
}

func init() {
	rootCmd.AddCommand(displayProductCmd)
}
