package consum

import (
	"errors"
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/consum"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

var iconNote string

func init() {
	cmd.RootCmd.AddCommand(NewInsertConsumptionsCmd())
}

func NewInsertConsumptionsCmd() *cobra.Command {
	command := newInsertConsumptionsCmd(consum.NewConsumptionsManager())
	command.Flags().StringVarP(&iconNote, "nota", "n", "", "Afegeix una nota al consum")
	return command
}

func newInsertConsumptionsCmd(manager consum.ConsumptionsManager) *cobra.Command {
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
		Args: util.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ica, err := parseInsertConsumptionsArgs(args, iconNote)
			if err != nil {
				return err
			}

			msg, err := manager.InsertConsumptions(ica)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

func parseInsertConsumptionsArgs(args []string, noteArg string) (consum.InsertConsumptionsArgs, error) {
	code, err := util.ParseInteger(args[0], "d'infant")
	if err != nil {
		return consum.InsertConsumptionsArgs{}, err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return consum.InsertConsumptionsArgs{}, errors.New("No s'ha indroduit el codi del darrer producte")
		}

		consUnits, err := util.ParseFloat(args[i])
		if err != nil {
			return consum.InsertConsumptionsArgs{}, err
		}

		productCode, err := util.ParseProductCode(args[i+1])
		if err != nil {
			return consum.InsertConsumptionsArgs{}, err
		}

		if _, ok := consMap[productCode]; ok {
			return consum.InsertConsumptionsArgs{}, errors.New("Hi ha un codi de producte repetit")
		}

		consMap[productCode] = consUnits
	}

	return consum.InsertConsumptionsArgs{Code: code, Consumptions: consMap, Note: noteArg}, nil
}
