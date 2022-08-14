package billing

import (
	"errors"
	"github.com/pjover/sam/internal/adapters/cli"
)

func ParseConsumptionsArgs(args []string, noteArg string) (id int, consumptions map[string]float64, note string, err error) {
	id, err = cli.ParseInteger(args[0], "d'infant")
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

		productId, err := cli.ParseProductId(args[i+1])
		if err != nil {
			return 0, nil, "", err
		}

		if _, ok := consMap[productId]; ok {
			return 0, nil, "", errors.New("hi ha un codi de producte repetit")
		}

		consMap[productId] = consUnits
	}

	return id, consMap, noteArg, nil
}
