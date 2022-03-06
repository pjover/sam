package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type displayProductCmd struct {
	displayService ports.DisplayService
}

func NewDisplayProductCmd(displayService ports.DisplayService) cli.Cmd {
	return displayProductCmd{
		displayService: displayService,
	}
}

func (e displayProductCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "mostraProducte codiProducte",
		Short:       "Mostra les dades d'un producte",
		Long:        "Mostra les dades d'un producte indicant el seu codi",
		Example:     `   mostraProducte age     Mostra les dades del producte AGE`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"mpro",
			"mostraproducte",
			"mostra-producte",
			"mostrarProducte",
			"mostrarproducte",
			"mostrar-producte",
			"dpro",
			"displayProduct",
			"displayproduct",
			"display-product",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cli.ParseProductId(args[0])
			if err != nil {
				return err
			}

			msg, err := e.displayService.DisplayProduct(id)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}
