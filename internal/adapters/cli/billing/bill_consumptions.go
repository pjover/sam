package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type billConsumptionsCmd struct {
	service ports.BillingService
}

func NewBillConsumptionsCmd(service ports.BillingService) cli.Cmd {
	return billConsumptionsCmd{
		service: service,
	}
}

func (i billConsumptionsCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "facturaConsums",
		Short:       "Factura els consums no facturats",
		Long:        `Factura els consums pendents de facturar de tots els infants i mostra el resum de les factures`,
		Example:     `   facturaConsums     Factura els consums no facturats`,
		Annotations: map[string]string{"CON": "Comandes de consum"},
		Aliases: []string{
			"fcon",
			"facturaconsums",
			"factura-consums",
			"facturarConsums",
			"facturarconsums",
			"facturar-consums",
			"bcon",
			"billConsumptions",
			"billconsumptions",
			"bill-consum",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := i.service.BillConsumptions()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
