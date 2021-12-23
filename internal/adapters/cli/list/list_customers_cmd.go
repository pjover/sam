package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/cobra"
)

type listCustomersCmd struct {
	listService ports.ListService
}

func NewListCustomersCmd(listService ports.ListService) cli.Cmd {
	return listCustomersCmd{
		listService: listService,
	}
}

func (l listCustomersCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "llistaClients",
		Short:       "Llista tots els clients",
		Long:        "Llista tots els clients i els seus infants per grups",
		Example:     `   llistaClients      Llista tots els clients`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"lcli",
			"llistaclients",
			"llista-clients",
			"llistarClients",
			"llistarclients",
			"llistar-clients",
			"lcus",
			"listCustomers",
			"listcustomers",
			"list-customers",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := l.listService.ListCustomers()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
