package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/internal/cmd"
	"sam/internal/list"
)

var listProductsCmd = &cobra.Command{
	Use:         "llistaProductes",
	Short:       "Llista tots els productes",
	Long:        "Llista tots els productes que hi han a la base de dades",
	Example:     `   llistaProductes    "Llista tots els productes`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"lpro",
		"llistaproductes", "llista-productes",
		"llistarProductes", "llistarproductes", "llistar-productes",
		"listProducts", "listproducts", "list-products",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := list.NewListManager()
		msg, err := manager.ListProducts()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(listProductsCmd)
}
