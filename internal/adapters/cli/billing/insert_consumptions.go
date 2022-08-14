package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

var iconNote string

type insertConsumptionsCmd struct {
	service ports.BillingService
}

func NewInsertConsumptionsCmd(service ports.BillingService) cli.Cmd {
	return insertConsumptionsCmd{
		service: service,
	}
}

func (i insertConsumptionsCmd) Cmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "insertaConsums codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota]",
		Short: "Inserta consums per a un infant",
		Long: `Inserta consums per a un infant al mes de treball
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
		Example: `   insertaConsums 2460 1 qme 0.5 mme      Inserta un consum per l'infant 2460 d'un QME i mig MME
   insertaConsums 2460 1 QME -n "Nota"    Inserta un consum per l'infant 2460 d'un QME amb una nota
   insertaConsums 2460 -- -5 GEN          Inserta un consum negatiu de -5 GEN (per posar un número negatiu s'han de possar dos guionets abans i no se poden posar notes)`,
		Annotations: map[string]string{"CON": "Comandes de consum"},
		Aliases: []string{
			"icon",
			"insertaconsums",
			"inserta-consums",
			"insertarConsums",
			"insertarconsums",
			"insertar-consums",
			"insertConsumptions",
			"insertconsumptions",
			"insert-consum",
		},
		Args: cli.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, consumptions, note, err := ParseConsumptionsArgs(args, iconNote)
			if err != nil {
				return err
			}

			msg, err := i.service.InsertConsumptions(id, consumptions, note)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
	command.Flags().StringVarP(&iconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}
