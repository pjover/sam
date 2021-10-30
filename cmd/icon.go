package cmd

import (
	"github.com/spf13/cobra"
	"sam/cons"
)

var note bool

var iconCmd = &cobra.Command{
	Use:   "icon codiInfant unitats codiProducte [-n nota]",
	Short: "Inserta consums per a un infant",
	Long: `Inserta consums per a un infant al mes de treball
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
	Example: `   icon 1520 1 QME 0.5 MME      Inserta un consum per l'infant 1520 d'un QME i mig MME
   icon 1520 1 QME -n "Això és una nota"    Inserta un consum per l'infant 1520 d'un QME amb una nota`,
	Annotations: map[string]string{"CON": "Comandes de consum"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ica, err := parseInsertConsumptionsArgs(args)
		if err != nil {
			return err
		}
		return cons.InsertConsumptions(ica)
	},
}

func init() {
	rootCmd.AddCommand(iconCmd)
	iconCmd.Flags().BoolVarP(&note, "nota", "n", false, "Afegeix una nota al consum")
}

func parseInsertConsumptionsArgs(args []string) (cons.InsertConsumptionsArgs, error) {
	ica := cons.InsertConsumptionsArgs{}
	return ica, nil
}
