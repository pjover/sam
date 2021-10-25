package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var bcliCmd = &cobra.Command{
	Use:         "bcli nomDelClient",
	Short:       "Busca els clients que tenguin 'nomDelClient'",
	Long:        "Busca els clients que tenguin 'nomDelClient' al camps de texte",
	Example:     `   mcli maria     Mostra les dades del clients amb 'maria'`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateArgsExists(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return adm.SearchCustomer(args)
	},
}

func init() {
	rootCmd.AddCommand(bcliCmd)
}
