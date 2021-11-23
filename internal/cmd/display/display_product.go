package display

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/display"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(NewDisplayProductCmd())
}

func NewDisplayProductCmd() *cobra.Command {
	return newDisplayProductCmd(display.NewProductDisplay())
}

func newDisplayProductCmd(dsp display.Display) *cobra.Command {
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
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			code, err := util.ParseProductCode(args[0])
			if err != nil {
				return err
			}

			msg, err := dsp.Display(code)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
