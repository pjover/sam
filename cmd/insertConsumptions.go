package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"sam/core"
)

var iconNote string

var insertConsumptionsCmd = &cobra.Command{
	Use:   "insertaConsums codiInfant unitats codiProducte [unitats codiProducte ...] [-n nota]",
	Short: "Inserta consums per a un infant",
	Long: `Inserta consums per a un infant al mes de treball
   - El mes de treball és el determinat per a l'execució de la comanda dir`,
	Example: `   insertaConsums 1520 1 qme 0.5 mme      Inserta un consum per l'infant 1520 d'un QME i mig MME
   insertaConsums 1520 1 QME -n "Nota"    Inserta un consum per l'infant 1520 d'un QME amb una nota
   insertaConsums 1520 -- -5 GEN          Inserta un consum negatiu de -5 GEN (per posar un número negatiu s'han de possar dos guionets abans i no se poden posar notes)`,
	Annotations: map[string]string{"CON": "Comandes de consum"},
	Aliases: []string{
		"icon",
		"insertaconsums", "inserta-consums",
		"insertarConsums", "insertarconsums", "insertar-consums",
		"insertConsumptions", "insertconsumptions", "insert-consumptions",
	},
	Args: MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		ica, err := parseInsertConsumptionsArgs(args, iconNote)
		if err != nil {
			return err
		}

		manager := core.NewConsumptionsManager()
		_, err = manager.InsertConsumptions(ica)
		return err
	},
}

func init() {
	rootCmd.AddCommand(insertConsumptionsCmd)
	insertConsumptionsCmd.Flags().StringVarP(&iconNote, "nota", "n", "", "Afegeix una nota al consum")
}

func parseInsertConsumptionsArgs(args []string, noteArg string) (core.InsertConsumptionsArgs, error) {
	code, err := parseInteger(args[0], "d'infant")
	if err != nil {
		return core.InsertConsumptionsArgs{}, err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return core.InsertConsumptionsArgs{}, errors.New("No s'ha indroduit el codi del darrer producte")
		}

		consUnits, err := parseFloat(args[i])
		if err != nil {
			return core.InsertConsumptionsArgs{}, err
		}

		productCode, err := parseProductCode(args[i+1])
		if err != nil {
			return core.InsertConsumptionsArgs{}, err
		}

		if _, ok := consMap[productCode]; ok {
			return core.InsertConsumptionsArgs{}, errors.New("Hi ha un codi de producte repetit")
		}

		consMap[productCode] = consUnits
	}

	return core.InsertConsumptionsArgs{Code: code, Consumptions: consMap, Note: noteArg}, nil
}
