package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/os"

	"github.com/pjover/sam/internal/list"
	"github.com/spf13/cobra"
)

type listInvoicesCmd struct {
	listService list.ListInvoices
	timeManager os.TimeManager
}

func NewListInvoicesCmd() *cobra.Command { //TODO Acabar d'adaptar a hex arch
	l := listInvoicesCmd{
		list.NewListInvoices(),
		os.NewTimeManager(),
	}
	return l.Cmd()
}

func (l listInvoicesCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "llistaFactures [codiClient] [AAAA-MM]",
		Short: "Llista les factura del mes i del client",
		Long: `Llista les factura del mes i del client
    - si no s'especifica el mes agafa l'actual
    - si no s'especifica client, llista les factures de tots els clients'`,
		Example: `   llistaFactures               Llista les factura del mes actual
   listaFactures 2021-10        Llista totes les factura del mes d'Octubre de 2021
   listaFactures 222            Llista totes les factura del client 222
   listaFactures 222 2021-10    Llista les factura del mes d'Octubre de 2021 del client 222`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"lfac",
			"llistafactures",
			"llista-factures",
			"llistarFactures",
			"llistarfactures",
			"llistar-factures",
			"linv",
			"listInvoices",
			"listinvoices",
			"list-invoices",
		},
		Args: cli.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return l.parseListInvoicesArgs(args)
		},
	}
}

func (l listInvoicesCmd) parseListInvoicesArgs(args []string) error {
	var msg string
	var err error
	switch len(args) {
	case 0:
		var workingTime = l.timeManager.Now()
		msg, err = l.listService.ListYearMonthInvoices(workingTime)
	case 1:
		customerCode, err := cli.ParseInteger(args[0], "de client")
		if err == nil {
			msg, err = l.listService.ListCustomerInvoices(customerCode)
		}
		yearMonth, err := cli.ParseYearMonth(args[0])
		if err == nil {
			msg, err = l.listService.ListYearMonthInvoices(yearMonth)
		}
	case 2:
		customerCode, err := cli.ParseInteger(args[0], "de client")
		if err != nil {
			return err
		}
		yearMonth, err := cli.ParseYearMonth(args[1])
		if err != nil {
			return err
		}
		msg, err = l.listService.ListCustomerYearMonthInvoices(customerCode, yearMonth)
	}
	if err != nil {
		return err
	}

	fmt.Println(msg)
	return nil
}
