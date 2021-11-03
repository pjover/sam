package cmd

import (
	"sam/adm"
	"strconv"

	"github.com/spf13/cobra"
)

var lconCmd = &cobra.Command{
	Use:         "lcon codiClient",
	Short:       "Mostra els consums d'un client",
	Long:        "Mostra el consums no facturats d'un client, indicant el seu codi",
	Example:     `   lcon 152     Mostra el consums del client 152`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-consums"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		return validateCustomerCode(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		if len(args) != 0 {
			customerCode, _ := strconv.Atoi(args[0])
			_, err := manager.ListCustomerConsumptions(customerCode)
			return err
		} else {
			_, err := manager.ListAllCustomersConsumptions()
			return err
		}
	},
}

func init() {
	rootCmd.AddCommand(lconCmd)
}
