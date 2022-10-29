package billing

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

var rconMonth string
var rconNote string

type rectifyConsumptionsCmd struct {
	configService ports.ConfigService
	service       ports.BillingService
}

func NewRectifyConsumptionsCmd(configService ports.ConfigService, service ports.BillingService) cli.Cmd {
	return rectifyConsumptionsCmd{
		configService: configService,
		service:       service,
	}
}

func (i rectifyConsumptionsCmd) Cmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "rectificaConsums codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota] [-m mes]",
		Short: "Rectifica els consums d'un infant",
		Long: `Rectifica els consums d'un infant al mes indicat o si no al mes de treball
   - Les rectificacions serveixen per corregir consums ja facturats, i se facturen amb la serie R`,
		Example: `   rectificaConsum 2460 1 QME 0.5 MME       Rectifica un consum per l'infant 2460 d'un QME i mig MME
   rectificaConsums 2460 1 QME -n "Nota"    Rectifica un consum per l'infant 2460 d'un QME amb una nota
   rectificaConsums 2460 1 QME -m 2022-10   Rectifica un consum al mes d'octubre de 2022 per l'infant 2460 d'un QME
   rectificaConsums 2460 -- -5 GEN          Rectifica un consum negatiu de -5 GEN (per posar un número negatiu
s'han de possar dos guionets abans i no se poden posar notes ni altra mes que l'actual)`,
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
			id, consumptions, yearMonth, note, err := ParseConsumptionsArgs(args, rconMonth, rconNote)
			if err != nil {
				return err
			}

			msg, err := i.service.RectifyConsumptions(id, consumptions, yearMonth, note)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
	var defaultMonth = i.configService.GetCurrentYearMonth().String()
	command.Flags().StringVarP(&rconMonth, "mes", "m", defaultMonth, "Defineix el mes de consum (en format 2022-10")
	command.Flags().StringVarP(&rconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}
