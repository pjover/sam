package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"sam/adm"
	"sam/util"
)

var listInvoicesCmd = &cobra.Command{
	Use:   "llistaFactures [codiClient] [AAAA-MM]",
	Short: "Llista les factura del mes i del client",
	Long: `Llista les factura del mes i del client,
    - si no s'especifica el mes agafa l'actual
    - si no s'especifica client, llista les factures de tots els clients'`,
	Example: `   lfac                         Llista les factura del mes actual
   listaFactures 2021-10        Llista totes les factura del mes d'Octubre de 2021
   listaFactures 222            Llista totes les factura del client 222
   listaFactures 222 2021-10    Llista les factura del mes d'Octubre de 2021 del client 222`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"lfac",
		"llistafactures", "llista-factures",
		"llistarFactures", "llistarfactures", "llistar-factures",
		"linv",
		"listInvoices", "listinvoices", "list-invoices",
	},
	Args: RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return parseListInvoicesArgs(args)
	},
}

func init() {
	rootCmd.AddCommand(listInvoicesCmd)
}

func parseListInvoicesArgs(args []string) error {
	manager := adm.NewListManager()
	switch len(args) {
	case 0:
		var workingTime = util.SamTimeManager{}.Now()
		_, err := manager.ListYearMonthInvoices(workingTime)
		return err
	case 1:
		customerCode, err := parseInteger(args[0], "de client")
		if err == nil {
			_, err := manager.ListCustomerInvoices(customerCode)
			return err
		}
		yearMonth, err := parseYearMonth(args[0])
		if err == nil {
			_, err := manager.ListYearMonthInvoices(yearMonth)
			return err
		}
		return err
	case 2:
		customerCode, err := parseInteger(args[0], "de client")
		if err != nil {
			return err
		}
		yearMonth, err := parseYearMonth(args[1])
		if err != nil {
			return err
		}
		_, err = manager.ListCustomerYearMonthInvoices(customerCode, yearMonth)
		return err
	}
	return errors.New("Unknown error")
}
