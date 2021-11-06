package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var listConsumptionsCmd = &cobra.Command{
	Use:   "llistaConsums [codiClient]",
	Short: "Mostra els consums d'un client",
	Long:  "Mostra el consums no facturats d'un client, indicant el seu codi",
	Example: `   llistaConsums 152     Mostra el consums del client 152
   llistaConsums         Mostra el consums de tots els clients`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"listConsumptions", "lcon",
		"llista-consums", "llistaconsums",
		"llistar-consums", "llistarconsums", "llistarConsums",
	},
	Args: RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		if len(args) != 0 {
			customerCode, err := parseInteger(args[0], "de client")
			if err != nil {
				return err
			}
			_, err = manager.ListCustomerConsumptions(customerCode)
			return err
		} else {
			_, err := manager.ListAllCustomersConsumptions()
			return err
		}
	},
}

func init() {
	rootCmd.AddCommand(listConsumptionsCmd)
}
