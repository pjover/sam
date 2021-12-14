package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/cobra"
)

type editProductCmd struct {
	editService ports.EditService
}

func NewEditProductCmd(editService ports.EditService) cli.Cmd {
	return editProductCmd{
		editService: editService,
	}
}

func (e editProductCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "editaProducte codiProducte",
		Short:       "Edita les dades d'un producte",
		Long:        "Obri un navegador per a editar les dades d'un producte indicant el seu codi",
		Example:     `   editaProducte age     Edita les dades del producte AGE`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"epro",
			"editaproducte",
			"edita-producte",
			"editarProducte",
			"editarproducte",
			"editarproducte",
			"editProduct",
			"editproduct",
			"edit-product",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			code, err := cli.ParseProductCode(args[0])
			if err != nil {
				return err
			}
			msg, err := e.editService.EditProduct(code)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}
