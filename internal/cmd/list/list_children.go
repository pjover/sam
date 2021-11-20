package list

import (
	"fmt"
	"sam/internal/cmd"
	"sam/internal/list"

	"github.com/spf13/cobra"
)

func newListChildrenCmd(manager list.ListChildren) *cobra.Command {
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
			msg, err := manager.List()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

func init() {
	manager := list.NewListChildren()
	command := newListChildrenCmd(manager)
	cmd.RootCmd.AddCommand(command)
}
