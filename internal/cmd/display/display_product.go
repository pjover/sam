package display

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/internal/cmd"
	"sam/internal/display"
	"sam/internal/util"
)

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

func init() {
	dsp := display.NewProductDisplay()
	command := newDisplayProductCmd(dsp)
	cmd.RootCmd.AddCommand(command)
}
