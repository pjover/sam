package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/cobra"
)

type listProductsCmd struct {
	listService ports.ListService
}

func NewListProductsCmd(listService ports.ListService) cli.Cmd {
	return listProductsCmd{
		listService: listService,
	}
}

func (e listProductsCmd) Cmd() *cobra.Command {
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
			msg, err := e.listService.ListProducts()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
