package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"sam/cons"
	"strings"
)

var note string

var iconCmd = &cobra.Command{
	Use:   "icon codiInfant unitats codiProducte [-n nota]",
	Short: "Inserta consums per a un infant",
	Long: `Inserta consums per a un infant al mes de treball
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
	Example: `   icon 1520 1 QME 0.5 MME      Inserta un consum per l'infant 1520 d'un QME i mig MME
   icon 1520 1 QME -n "Això és una nota"    Inserta un consum per l'infant 1520 d'un QME amb una nota`,
	Annotations: map[string]string{"CON": "Comandes de consum"},
	Aliases:     []string{"inserta-consum"},
	RunE: func(cmd *cobra.Command, args []string) error {
		ica, err := parseInsertConsumptionsArgs(args, note)
		if err != nil {
			return err
		}
		return cons.InsertConsumptions(ica)
	},
}

func init() {
	rootCmd.AddCommand(iconCmd)
	iconCmd.Flags().StringVarP(&note, "nota", "n", "", "Afegeix una nota al consum")
}

func parseInsertConsumptionsArgs(args []string, noteArg string) (cons.InsertConsumptionsArgs, error) {
	err := validateNumberOfArgsGreaterThan(3, args)
	if err != nil {
		return cons.InsertConsumptionsArgs{}, err
	}

	code, err := parseInteger(args[0])
	if err != nil {
		return cons.InsertConsumptionsArgs{}, err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return cons.InsertConsumptionsArgs{}, errors.New("No s'ha indroduit el codi del darrer consum")
		}
		consUnits, err := parseFloat(args[i])
		if err != nil {
			return cons.InsertConsumptionsArgs{}, err
		}
		consCode := strings.ToUpper(args[i+1])
		consMap[consCode] = consUnits
	}

	return cons.InsertConsumptionsArgs{Code: code, Consumptions: consMap, Note: noteArg}, nil
}
