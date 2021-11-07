package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var displayCustomerCmd = &cobra.Command{
	Use:         "mostraClient codiClient",
	Short:       "Mostra les dades d'un client",
	Long:        "Mostra les dades d'un client indicant el seu codi",
	Example:     `   mostraClient 152     Mostra les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"mcli",
		"mostraclient", "mostra-client",
		"mostrarClient", "mostrarclient", "mostrar-client",
		"dcus",
		"displayCustomer", "displaycustomer", "display-customer",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		customerCode, _ := parseInteger(args[0], "de client")
		manager := adm.NewDisplayManager()
		_, err := manager.DisplayCustomer(customerCode)
		return err
	},
}

func init() {
	rootCmd.AddCommand(displayCustomerCmd)
}
