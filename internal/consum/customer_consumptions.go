package consum

import (
	"errors"
	"github.com/pjover/sam/internal/adapters/cli"
)

type CustomerConsumptionsManager interface {
	Run(args CustomerConsumptionsArgs) (string, error)
}

type CustomerConsumptionsArgs struct {
	Id           int
	Consumptions map[string]float64
	Note         string
}

func ParseInsertConsumptionsArgs(args []string, noteArg string) (CustomerConsumptionsArgs, error) {
	id, err := cli.ParseInteger(args[0], "d'infant")
	if err != nil {
		return CustomerConsumptionsArgs{}, err
	}

	var consMap = make(map[string]float64)
	for i := 1; i < len(args); i = i + 2 {
		if i >= len(args)-1 {
			return CustomerConsumptionsArgs{}, errors.New("no s'ha indroduit el codi del darrer producte")
		}

		consUnits, err := cli.ParseFloat(args[i])
		if err != nil {
			return CustomerConsumptionsArgs{}, err
		}

		productId, err := cli.ParseProductId(args[i+1])
		if err != nil {
			return CustomerConsumptionsArgs{}, err
		}

		if _, ok := consMap[productId]; ok {
			return CustomerConsumptionsArgs{}, errors.New("hi ha un codi de producte repetit")
		}

		consMap[productId] = consUnits
	}

	return CustomerConsumptionsArgs{Id: id, Consumptions: consMap, Note: noteArg}, nil
}
