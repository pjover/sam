package list

import (
	"fmt"
	"sam/adm"
	"sam/internal/cmd"

	"github.com/spf13/cobra"
)

var listCustomersCmd = &cobra.Command{
	Use:         "llistaClients",
	Short:       "Llista tots els clients",
	Long:        "Llista tots els clients i els seus infants per grups",
	Example:     `   llistaClients      Llista tots els clients`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"lcli",
		"llistaclients", "llista-clients",
		"llistarClients", "llistarclients", "llistar-clients",
		"lcus",
		"listCustomers", "listcustomers", "list-customers",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		msg, err := manager.ListCustomers()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(listCustomersCmd)
}
