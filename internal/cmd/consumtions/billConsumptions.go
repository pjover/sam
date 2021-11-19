package consumtions

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/core"
	"sam/internal/cmd"
)

var billConsumptionsCmd = &cobra.Command{
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
		"facturaconsums", "factura-consums",
		"facturarConsums", "facturarconsums", "facturar-consums",
		"bcon",
		"billConsumptions", "billconsumptions", "bill-consumptions",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := core.NewConsumptionsManager()
		msg, err := manager.BillConsumptions()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(billConsumptionsCmd)
}
