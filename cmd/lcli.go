package cmd

import (
	"sam/adm"

	"github.com/spf13/cobra"
)

var lcliCmd = &cobra.Command{
	Use:         "lcli",
	Short:       "Llista tots els clients",
	Long:        "Llista tots els clients i els seus infants per grups",
	Example:     `   lcli      Llista tots els clients`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-clients"},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		_, err := manager.ListCustomers()
		return err
	},
}

func init() {
	rootCmd.AddCommand(lcliCmd)
}
