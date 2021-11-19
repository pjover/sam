package cmd

import (
	"fmt"
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
		msg, err := manager.SearchCustomer(args)
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(searchCustomerCmd)
}
