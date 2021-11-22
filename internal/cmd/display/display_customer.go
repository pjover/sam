package display

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/display"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

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
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := util.ParseInteger(args[0], "de client")
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

func init() {
	dsp := display.NewCustomerDisplay()
	command := newDisplayCustomerCmd(dsp)
	cmd.RootCmd.AddCommand(command)
}
