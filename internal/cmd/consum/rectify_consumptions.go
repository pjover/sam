package consum

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/internal/cmd"
	"sam/internal/consum"
	"sam/internal/util"
)

var rconNote string

func newRectifyConsumptionsCmd(manager consum.ConsumptionsManager) *cobra.Command {
	return &cobra.Command{
		Use:   "rectificaConsums codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota]",
		Short: "Rectifica els consums d'un infant",
		Long: `Rectifica els consums d'un infant al mes de treball
   - Les rectificacions serveixen per corretgir consums ja facturats, i se facturen amb la serie R
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
		Example: `   rectificaConsum 2460 1 QME 0.5 MME       Rectifica un consum per l'infant 2460 d'un QME i mig MME
   rectificaConsums 2460 1 QME -n "Nota"    Rectifica un consum per l'infant 2460 d'un QME amb una nota
   rectificaConsums 2460 -- -5 GEN          Rectifica un consum negatiu de -5 GEN (per posar un número negatiu s'han de possar dos guionets abans i no se poden posar notes)`,
		Annotations: map[string]string{"CON": "Comandes de consum"},
		Aliases: []string{
			"rcon",
			"rectificaconsums",
			"rectifica-consums",
			"rectificarConsums",
			"rectificarconsums",
			"rectificar-consums",
			"rectifyConsumptions",
			"rectifyconsumptions",
			"rectify-consum",
		},
		Args: util.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ica, err := parseInsertConsumptionsArgs(args, rconNote)
			if err != nil {
				return err
			}

			msg, err := manager.RectifyConsumptions(ica)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

func init() {
	manager := consum.NewConsumptionsManager()
	command := newRectifyConsumptionsCmd(manager)
	command.Flags().StringVarP(&rconNote, "nota", "n", "", "Afegeix una nota al consum")
	cmd.RootCmd.AddCommand(command)
}
