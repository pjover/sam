package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/spf13/cobra"
)

type billConsumptionsCmd struct {
	service billing.BillingService
}

func NewBillConsumptionsCmd(service billing.BillingService) cli.Cmd {
	return billConsumptionsCmd{
		service: service,
	}
}

func (i billConsumptionsCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "facturaConsums",
		Short: "Factura els consums no facturats",
		Long: `Factura els consums pendents de facturar de tots els infants
   - Mostra el resum de les factures
   - Crea els PDFs de les factures al directori 'factures' dedins del directori de treball
   - Crea el fitxer de rebuts 'bdd-n.q1x' de les factures de rebuts dedins del directori de treball
     - 'n' és el número de sequència, per si hi ha més d'un fitxer
   - Crea el fitxer resum de factures 'factures.xlsx' dedins del directori de treball`,
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
