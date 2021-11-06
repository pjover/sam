package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var editCustomerCmd = &cobra.Command{
	Use:         "editaClient codiClient",
	Short:       "Edita les dades d'un client",
	Long:        "Obri un navegador per a editar les dades d'un client indicant el seu codi",
	Example:     `   editaClient 152     Edita les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"editCustomer", "ecli",
		"editaclient", "edita-client",
		"editar-client", "editarclient", "editarClient",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		customerCode, err := parseInteger(args[0], "de client")
		if err != nil {
			return err
		}
		return adm.EditCustomer(customerCode)
	},
}

func init() {
	rootCmd.AddCommand(editCustomerCmd)
}
