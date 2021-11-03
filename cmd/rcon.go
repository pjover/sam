package cmd

import (
	"github.com/spf13/cobra"
	"sam/core"
)

var rconNote string

var rconCmd = &cobra.Command{
	Use:   "rcon codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota]",
	Short: "Rectifica els consums d'un infant",
	Long: `Rectifica els consums d'un infant al mes de treball
   - Les rectificacions serveixen per corretgir consums ja facturats, i se facturen amb la serie R
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
	Example: `   rcon 1520 1 QME 0.5 MME      Rectifica un consum per l'infant 1520 d'un QME i mig MME
   rcon 1520 1 QME -n "Això és una nota"    Rectifica un consum per l'infant 1520 d'un QME amb una nota
   rcon 1520 -- -5 GEN    Rectifica un consum negatiu de -5 GEN (per posar un número negatiu s'han de possar dos guionets abans i no se poden posar notes)`,
	Annotations: map[string]string{"CON": "Comandes de consum"},
	Aliases:     []string{"rectifica-consum"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ica, err := parseInsertConsumptionsArgs(args, rconNote)
		if err != nil {
			return err
		}

		manager := core.NewConsumptionsManager()
		_, err = manager.RectifyConsumptions(ica)
		return err
	},
}

func init() {
	rootCmd.AddCommand(rconCmd)
	rconCmd.Flags().StringVarP(&rconNote, "nota", "n", "", "Afegeix una nota al consum")
}
