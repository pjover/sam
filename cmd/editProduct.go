package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var editProductCmd = &cobra.Command{
	Use:         "editaProducte codiProducte",
	Short:       "Edita les dades d'un producte",
	Long:        "Obri un navegador per a editar les dades d'un producte indicant el seu codi",
	Example:     `   editaProducte age     Edita les dades del producte AGE`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"epro",
		"editaproducte", "edita-producte",
		"editarProducte", "editarproducte", "editarproducte",
		"editProduct", "editproduct", "edit-product",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		productCode, err := parseProductCode(args[0])
		if err != nil {
			return err
		}
		return adm.EditProduct(productCode)
	},
}

func init() {
	RootCmd.AddCommand(editProductCmd)
}
