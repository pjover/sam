package cmd

import (
	"fmt"
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
	Example: `   llistaFactures               Llista les factura del mes actual
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
	var msg string
	var err error
	switch len(args) {
	case 0:
		var workingTime = util.SamTimeManager{}.Now()
		msg, err = manager.ListYearMonthInvoices(workingTime)
	case 1:
		customerCode, err := parseInteger(args[0], "de client")
		if err == nil {
			msg, err = manager.ListCustomerInvoices(customerCode)
		}
		yearMonth, err := parseYearMonth(args[0])
		if err == nil {
			msg, err = manager.ListYearMonthInvoices(yearMonth)
		}
	case 2:
		customerCode, err := parseInteger(args[0], "de client")
		if err != nil {
			return err
		}
		yearMonth, err := parseYearMonth(args[1])
		if err != nil {
			return err
		}
		msg, err = manager.ListCustomerYearMonthInvoices(customerCode, yearMonth)
	}
	if err != nil {
		return err
	}

	fmt.Println(msg)
	return nil
}
