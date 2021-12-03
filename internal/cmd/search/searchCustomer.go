package search

import (
	"fmt"

	"github.com/pjover/sam/internal/search"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
)

func NewSearchCustomerCmd() *cobra.Command {
	return newSearchCustomerCmd(search.NewSearchManager())
}

func newSearchCustomerCmd(manager search.SearchManager) *cobra.Command {
	return &cobra.Command{
		Use:         "buscaClient nomDelClient",
		Short:       "Busca els clients que tenguin 'nomDelClient'",
		Long:        "Busca els clients que tenguin 'nomDelClient' al camps de texte",
		Example:     `   buscaClient maria     Mostra les dades dels clients amb 'maria'`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"bcli",
			"buscaclient",
			"busca-client",
			"buscarClient",
			"buscarclient",
			"buscar-client",
			"buscarClients",
			"buscarclients",
			"buscar-clients",
			"scus",
			"searchCustomer",
			"searchcustomer",
			"search-customer",
		},
		Args: shared.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := manager.SearchCustomer(args)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
