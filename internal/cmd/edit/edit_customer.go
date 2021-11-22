package edit

import (
	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/edit"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

func newEditCustomerCmd(editor edit.Editor) *cobra.Command {
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
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := util.ParseInteger(args[0], "de client")
			if err != nil {
				return err
			}
			return editor.Edit(args[0])
		},
	}
}

func init() {
	editor := edit.NewCustomerEditor()
	command := newEditCustomerCmd(editor)
	cmd.RootCmd.AddCommand(command)
}
