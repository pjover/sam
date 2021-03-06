package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

var rconNote string

type rectifyConsumptionsCmd struct {
	service ports.BillingService
}

func NewRectifyConsumptionsCmd(service ports.BillingService) cli.Cmd {
	return rectifyConsumptionsCmd{
		service: service,
	}
}

func (i rectifyConsumptionsCmd) Cmd() *cobra.Command {
	command := &cobra.Command{
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
		Args: cli.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, consumptions, note, err := parseConsumptionsArgs(args, rconNote)
			if err != nil {
				return err
			}

			msg, err := i.service.RectifyConsumptions(id, consumptions, note)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
	command.Flags().StringVarP(&rconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}
