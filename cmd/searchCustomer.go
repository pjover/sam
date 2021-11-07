package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var searchCustomerCmd = &cobra.Command{
	Use:         "buscaClient nomDelClient",
	Short:       "Busca els clients que tenguin 'nomDelClient'",
	Long:        "Busca els clients que tenguin 'nomDelClient' al camps de texte",
	Example:     `   buscaClient maria     Mostra les dades dels clients amb 'maria'`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"bcli",
		"buscaclient", "busca-client",
		"buscarClient", "buscarclient", "buscar-client",
		"buscarClients", "buscarclients", "buscar-clients",
		"scus",
		"searchCustomer", "searchcustomer", "search-customer",
	},
	Args: MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewSearchManager()
		_, err := manager.SearchCustomer(args)
		return err
	},
}

func init() {
	rootCmd.AddCommand(searchCustomerCmd)
}
