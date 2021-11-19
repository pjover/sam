package display

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/adm"
	"sam/internal/cmd"
	"sam/internal/util"
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
	Args: util.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		customerCode, _ := util.ParseInteger(args[0], "de client")
		manager := adm.NewDisplayManager()
		msg, err := manager.DisplayCustomer(customerCode)
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(displayCustomerCmd)
}
