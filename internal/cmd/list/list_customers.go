package list

import (
	"fmt"

	"github.com/pjover/sam/internal/list"

	"github.com/spf13/cobra"
)

func NewListCustomersCmd() *cobra.Command {
	return newListCustomersCmd(list.NewListCustomers())
}

func newListCustomersCmd(manager list.List) *cobra.Command {
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
			msg, err := manager.List()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
