package consum

import (
	"fmt"

	"github.com/pjover/sam/internal/consum"
	"github.com/spf13/cobra"
)

func NewBillConsumptionsCmd() *cobra.Command {
	return newBillConsumptionsCmd(consum.NewBillConsumptionsManager())
}

func newBillConsumptionsCmd(manager consum.BillConsumptionsManager) *cobra.Command {
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
			msg, err := manager.Run()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}
