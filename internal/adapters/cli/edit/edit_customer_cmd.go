package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type editCustomerCmd struct {
	editService ports.EditService
}

func NewEditCustomerCmd(editService ports.EditService) cli.Cmd {
	return editCustomerCmd{
		editService: editService,
	}
}

func (e editCustomerCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "editaClient codiClient",
		Short:       "Edita les dades d'un client",
		Long:        "Obri un navegador per a editar les dades d'un client indicant el seu codi",
		Example:     `   editaClient 152     Edita les dades del client 152`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"ecli",
			"editaclient",
			"edita-client",
			"editarClient",
			"editarclient",
			"editar-client",
			"ecus",
			"editCustomer",
			"editcustomer",
			"edit-customer",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			code, err := cli.ParseInteger(args[0], "de client")
			if err != nil {
				return err
			}
			msg, err := e.editService.EditCustomer(code)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}
