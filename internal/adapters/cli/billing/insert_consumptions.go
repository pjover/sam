package billing

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/spf13/cobra"
)

var iconNote string

type insertConsumptionsCmd struct {
	service billing.BillingService
}

func NewInsertConsumptionsCmd(service billing.BillingService) cli.Cmd {
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
			code, consumptions, note, err := parseInsertConsumptionsArgs(args, iconNote)
			if err != nil {
				return err
			}

			msg, err := i.service.InsertConsumptions(code, consumptions, note)
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

func parseInsertConsumptionsArgs(args []string, noteArg string) (code int, consumptions map[string]float64, note string, err error) {
	code, err = cli.ParseInteger(args[0], "d'infant")
	if err != nil {
		return 0, nil, "", err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return 0, nil, "", errors.New("no s'ha indroduit el codi del darrer producte")
		}

		consUnits, err := cli.ParseFloat(args[i])
		if err != nil {
			return 0, nil, "", err
		}

		productCode, err := cli.ParseProductCode(args[i+1])
		if err != nil {
			return 0, nil, "", err
		}

		if _, ok := consMap[productCode]; ok {
			return 0, nil, "", errors.New("hi ha un codi de producte repetit")
		}

		consMap[productCode] = consUnits
	}

	return code, consMap, noteArg, nil
}
