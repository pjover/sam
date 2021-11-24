package list

import (
	"fmt"

	"github.com/pjover/sam/internal/list"
	"github.com/spf13/cobra"
)

func NewListProductsCmd() *cobra.Command {
	return newListProductsCmd(list.NewListProducts())
}

func newListProductsCmd(manager list.List) *cobra.Command {
	return &cobra.Command{
		Use:         "llistaProductes",
		Short:       "Llista tots els productes",
		Long:        "Llista tots els productes que hi han a la base de dades",
		Example:     `   llistaProductes    "Llista tots els productes`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"lpro",
			"llistaproductes",
			"llista-productes",
			"llistarProductes",
			"llistarproductes",
			"llistar-productes",
			"listProducts",
			"listproducts",
			"list-products",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := manager.List()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
