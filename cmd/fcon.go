package cmd

import (
	"github.com/spf13/cobra"
	"sam/core"
)

var fconCmd = &cobra.Command{
	Use:   "fcon",
	Short: "Factura els consums no facturats",
	Long: `Factura els consums pendents de facturar de tots els infants
   - Mostra el resum de les factures
   - Crea els PDFs de les factures al directori 'factures' dedins del directori de treball
   - Crea el fitxer de rebuts 'bdd-n.q1x' de les factures de rebuts dedins del directori de treball
     - 'n' és el número de sequència, per si hi ha més d'un fitxer
   - Crea el fitxer resum de factures 'factures.xlsx' dedins del directori de treball`,
	Example:     `   fcon     Factura els consums no facturats`,
	Annotations: map[string]string{"CON": "Comandes de consum"},
	Aliases:     []string{"factura-consums"},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := core.NewConsumptionsManager()
		return manager.BillConsumptions()
	},
}

func init() {
	rootCmd.AddCommand(fconCmd)
}
