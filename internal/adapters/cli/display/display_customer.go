package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type displayCustomerCmd struct {
	displayService ports.DisplayService
}

func NewDisplayCustomerCmd(displayService ports.DisplayService) cli.Cmd {
	return displayCustomerCmd{
		displayService: displayService,
	}
}

func (e displayCustomerCmd) Cmd() *cobra.Command {
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
			id, err := cli.ParseInteger(args[0], "de client")
			if err != nil {
				return err
			}
			msg, err := e.displayService.DisplayCustomer(id)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}
