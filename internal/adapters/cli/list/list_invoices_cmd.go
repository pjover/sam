package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"

	"github.com/spf13/cobra"
)

type listInvoicesCmd struct {
	configService ports.ConfigService
	listService   ports.ListService
}

func NewListInvoicesCmd(configService ports.ConfigService, listService ports.ListService) cli.Cmd {
	return listInvoicesCmd{
		configService: configService,
		listService:   listService,
	}
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
		yearMonth := l.configService.GetCurrentYearMonth()
		msg, err = l.listService.ListYearMonthInvoices(yearMonth)
	case 1:
		customerId, err := cli.ParseInteger(args[0], "de client")
		if err == nil {
			msg, err = l.listService.ListCustomerInvoices(customerId)
		}
		yearMonth, err := model.StringToYearMonth(args[0])
		if err == nil {
			msg, err = l.listService.ListYearMonthInvoices(yearMonth)
		}
	case 2:
		customerId, err := cli.ParseInteger(args[0], "de client")
		if err != nil {
			return err
		}
		yearMonth, err := model.StringToYearMonth(args[0])
		if err != nil {
			return err
		}
		msg, err = l.listService.ListCustomerYearMonthInvoices(customerId, yearMonth)
	}
	if err != nil {
		return err
	}

	fmt.Println(msg)
	return nil
}
