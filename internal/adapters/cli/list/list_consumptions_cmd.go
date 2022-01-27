package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"

	"github.com/spf13/cobra"
)

type listConsumptionsCmd struct {
	listService ports.ListService
}

func NewListConsumptionsCmd(listService ports.ListService) cli.Cmd {
	return listConsumptionsCmd{
		listService: listService,
	}
}

func (l listConsumptionsCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "llistaConsums [codiInfant]",
		Short: "Mostra els consums d'un infant",
		Long:  "Mostra el consums no facturats d'un infant, indicant el seu codi",
		Example: `   llistaConsums 2630    Mostra el consums de l'infant 2630
   llistaConsums         Mostra el consums de tots els clients`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"lcon",
			"llistaconsums",
			"llista-consums",
			"llistarConsums",
			"llistarconsums",
			"llistar-consums",
			"listConsumptions",
			"listconsumptions",
			"list-consum",
		},
		Args: cli.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var msg string
			var err error
			if len(args) != 0 {
				childCode, err := cli.ParseInteger(args[0], "de client")
				if err != nil {
					return err
				}
				msg, err = l.listService.ListChildConsumptions(childCode)
			} else {
				msg, err = l.listService.ListConsumptions()
			}
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
