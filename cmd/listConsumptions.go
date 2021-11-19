package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/adm"
	"sam/internal/util"
)

var listConsumptionsCmd = &cobra.Command{
	Use:   "llistaConsums [codiClient]",
	Short: "Mostra els consums d'un client",
	Long:  "Mostra el consums no facturats d'un client, indicant el seu codi",
	Example: `   llistaConsums 152     Mostra el consums del client 152
   llistaConsums         Mostra el consums de tots els clients`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"lcon",
		"llistaconsums", "llista-consums",
		"llistarConsums", "llistarconsums", "llistar-consums",
		"listConsumptions", "listconsumptions", "list-consumptions",
	},
	Args: util.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		var msg string
		var err error
		if len(args) != 0 {
			customerCode, err := util.ParseInteger(args[0], "de client")
			if err != nil {
				return err
			}
			msg, err = manager.ListCustomerConsumptions(customerCode)
		} else {
			msg, err = manager.ListAllCustomersConsumptions()
		}
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listConsumptionsCmd)
}
