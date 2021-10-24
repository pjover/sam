package cmd

import (
	"fmt"
	"sam/adm"

	"github.com/spf13/cobra"
)

var customerCode int
var childCode int

var mcliCmd = &cobra.Command{
	Use:   "mcli {-c 123 | -i 1230}",
	Short: "Mostra les dades d'un client",
	Long:  "Mostra les dades d'un client, se pot indicar el client o l'infant",
	Example: `   mcli -i 1520    Mostra les dades del client de l'infant 1520
   mcli -c 152     Mostra les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMcli()
	},
}

func init() {
	rootCmd.AddCommand(mcliCmd)
	mcliCmd.Flags().IntVarP(&customerCode, "client", "c", 0, "Codi del client")
	mcliCmd.Flags().IntVarP(&childCode, "infant", "i", 0, "Codi de l'infant")
}

func runMcli() error {
	if childCode > 0 && customerCode > 0 {
		return fmt.Errorf("Indicar el codi del client o del infant, no els dos a l'hora")
	} else if childCode == 0 && customerCode == 0 {
		return fmt.Errorf("És necesari indicar el codi del client o del infant")
	}

	if childCode > 0 {
		customerCode = childCode / 10
	}
	return adm.DisplayCustomer(customerCode)
}
