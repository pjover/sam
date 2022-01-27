package consum

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/ports"

	"github.com/pjover/sam/internal/consum"
	"github.com/spf13/cobra"
)

var iconNote string

func NewInsertConsumptionsCmd(httpPostManager hobbit.HttpPostManager, dbService ports.DbService) *cobra.Command {
	command := newInsertConsumptionsCmd(consum.NewInsertConsumptionsManager(httpPostManager, dbService))
	command.Flags().StringVarP(&iconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}

func newInsertConsumptionsCmd(manager consum.CustomerConsumptionsManager) *cobra.Command {
	return &cobra.Command{
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
			ica, err := parseInsertConsumptionsArgs(args, iconNote)
			if err != nil {
				return err
			}

			msg, err := manager.Run(ica)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

func parseInsertConsumptionsArgs(args []string, noteArg string) (consum.CustomerConsumptionsArgs, error) {
	code, err := cli.ParseInteger(args[0], "d'infant")
	if err != nil {
		return consum.CustomerConsumptionsArgs{}, err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return consum.CustomerConsumptionsArgs{}, errors.New("no s'ha indroduit el codi del darrer producte")
		}

		consUnits, err := cli.ParseFloat(args[i])
		if err != nil {
			return consum.CustomerConsumptionsArgs{}, err
		}

		productCode, err := cli.ParseProductCode(args[i+1])
		if err != nil {
			return consum.CustomerConsumptionsArgs{}, err
		}

		if _, ok := consMap[productCode]; ok {
			return consum.CustomerConsumptionsArgs{}, errors.New("hi ha un codi de producte repetit")
		}

		consMap[productCode] = consUnits
	}

	return consum.CustomerConsumptionsArgs{Code: code, Consumptions: consMap, Note: noteArg}, nil
}
