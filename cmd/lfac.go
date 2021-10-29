package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"sam/adm"
	"time"
)

var lfacCmd = &cobra.Command{
	Use:   "lfac [codiClient] [AAAA-MM]",
	Short: "Llista les factura del mes i del client",
	Long: `Llista les factura del mes i del client,
    - si no s'especifica el mes agafa l'actual
    - si no s'especifica client, llista les factures de tots els clients'`,
	Example: `   lfac                Llista les factura del mes actual
   lfac 2021-10        Llista totes les factura del mes d'Octubre de 2021
   lfac 222            Llista totes les factura del client 222
   lfac 222 2021-10    Llista les factura del mes d'Octubre de 2021 del client 222`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-factures"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgsBetween(0, 2, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return adm.ListYearMonthInvoices(time.Now())
		case 1:
			customerCode, err := parseIntegerCode(args[0])
			if err == nil {
				return adm.ListCustomerInvoices(customerCode)
			}
			yearMonth, err := parseYearMonth(args[0])
			if err == nil {
				return adm.ListYearMonthInvoices(yearMonth)
			}
			return err
		case 2:
			customerCode, err := parseIntegerCode(args[0])
			if err != nil {
				return err
			}
			yearMonth, err := parseYearMonth(args[1])
			if err != nil {
				return err
			}
			return adm.ListCustomerYearMonthInvoices(customerCode, yearMonth)
		}
		return errors.New("Unknown error")
	},
}

func init() {
	rootCmd.AddCommand(lfacCmd)
}
