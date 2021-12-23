package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/cobra"
)

type listChildrenCmd struct {
	listService ports.ListService
}

func NewListChildrenCmd(listService ports.ListService) cli.Cmd {
	return listChildrenCmd{
		listService: listService,
	}
}

func (l listChildrenCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "llistaInfants",
		Short:       "Llista tots els infants",
		Long:        "Llista tots els infants per grups",
		Example:     `   llistaInfants      Llista tots els infants`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"linf",
			"llistainfants", "llista-infants",
			"llistarInfants", "llistarinfants", "llistar-infants",
			"lchi",
			"listChildren", "listchildren", "list-children",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := l.listService.ListChildren()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}
