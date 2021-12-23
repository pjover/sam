package search

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/cobra"
)

type searchCustomerCmd struct {
	searchService ports.SearchService
}

func NewSearchCustomerCmd(searchService ports.SearchService) cli.Cmd {
	return searchCustomerCmd{
		searchService: searchService,
	}
}

func (e searchCustomerCmd) Cmd() *cobra.Command {
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
		Args: cli.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			searchText := fmt.Sprint(args)
			msg, err := e.searchService.SearchCustomer(searchText)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}
