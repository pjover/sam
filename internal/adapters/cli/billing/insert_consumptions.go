package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

var iconMonth string
var iconNote string

type insertConsumptionsCmd struct {
	configService ports.ConfigService
	service       ports.BillingService
}

func NewInsertConsumptionsCmd(configService ports.ConfigService, service ports.BillingService) cli.Cmd {
	return insertConsumptionsCmd{
		configService: configService,
		service:       service,
	}
}

func (i insertConsumptionsCmd) Cmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "insertaConsums codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota] [-m mes]",
		Short: "Inserta consums per a un infant",
		Long:  `Inserta consums per a un infant al mes indicat o si no al mes de treball`,
		Example: `   insertaConsums 2460 1 qme 0.5 mme      Inserta un consum per l'infant 2460 d'un QME i mig MME
   insertaConsums 2460 1 QME -n "Nota"    Inserta un consum al mes actual per l'infant 2460 d'un QME amb una nota
   insertaConsums 2460 1 QME -m 2022-10   Inserta un consum al mes d'octubre de 2022 per l'infant 2460 d'un QME
   insertaConsums 2460 -- -5 GEN          Inserta un consum negatiu de -5 GEN al mes actual (per posar un número negatiu
s'han de possar dos guionets abans i no se poden posar ni notes ni altra mes que l'actual)`,
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
			id, consumptions, yearMonth, note, err := ParseConsumptionsArgs(args, iconMonth, iconNote)
			if err != nil {
				return err
			}

			msg, err := i.service.InsertConsumptions(id, consumptions, yearMonth, note)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
	var defaultMonth = i.configService.GetCurrentYearMonth().String()
	command.Flags().StringVarP(&iconMonth, "mes", "m", defaultMonth, "Defineix el mes de consum (en format 2022-10")
	command.Flags().StringVarP(&iconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}
