package cmd

import (
	"sam/adm"

	"github.com/spf13/cobra"
)

var lcliCmd = &cobra.Command{
	Use:         "lcli codiClient",
	Short:       "Llista tots els clients",
	Long:        "Llista tots els clients i els infants per grups",
	Example:     `   lcli      Llista tots els clients`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"mostra-client"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return adm.ListCustomers()
	},
}

func init() {
	rootCmd.AddCommand(lcliCmd)
}
