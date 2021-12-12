package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/display"
	"github.com/spf13/cobra"
)

func NewDisplayCustomerCmd() *cobra.Command {
	return newDisplayCustomerCmd(display.NewCustomerDisplay())
}

func newDisplayCustomerCmd(dsp display.Display) *cobra.Command {
	return &cobra.Command{
		Use:         "mostraClient codiClient",
		Short:       "Mostra les dades d'un client",
		Long:        "Mostra les dades d'un client indicant el seu codi",
		Example:     `   mostraClient 152     Mostra les dades del client 152`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"mcli",
			"mostraclient",
			"mostra-client",
			"mostrarClient",
			"mostrarclient",
			"mostrar-client",
			"dcus",
			"displayCustomer",
			"displaycustomer",
			"display-customer",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := cli.ParseInteger(args[0], "de client")
			if err != nil {
				return err
			}

			msg, err := dsp.Display(args[0])
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
