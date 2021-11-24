package list

import (
	"fmt"

	"github.com/pjover/sam/internal/list"

	"github.com/spf13/cobra"
)

func NewListChildrenCmd() *cobra.Command {
	return newListChildrenCmd(list.NewListChildren())
}

func newListChildrenCmd(manager list.List) *cobra.Command {
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
