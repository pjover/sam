package cmd

import (
	"sam/adm"

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
		_, err := manager.ListCustomers()
		return err
	},
}

func init() {
	rootCmd.AddCommand(listCustomersCmd)
}
