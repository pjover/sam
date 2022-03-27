package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// RangeArgs returns an error if the number of args is not within the expected range.
func RangeArgs(min int, max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < min || len(args) > max {
			return fmt.Errorf("han d'esser entre %d i %d argument(s), rebuts %d", min, max, len(args))
		}
		return nil
	}
}

// MinimumNArgs returns an error if there is not at least N args.
func MinimumNArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < n {
			return fmt.Errorf("es requereixen al menys %d argument(s), rebuts %d", n, len(args))
		}
		return nil
	}
}

// ExactArgs returns an error if there are not exactly n args.
func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("es requereix %d argument(s), rebuts %d", n, len(args))
		}
		return nil
	}
}

func ParseInteger(strId string, idType string) (int, error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, fmt.Errorf("el codi %s introduit és invàlid: %s", idType, strId)
	}
	return id, nil
}

func ParseFloat(value string) (float64, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("el número introduit és invàlid: %s", value)
	}
	return float, nil
}

func ParseProductId(id string) (string, error) {
	if len(id) != 3 {
		return "", fmt.Errorf("el codi de producte introduit és invàlid: %s", id)
	}
	return strings.ToUpper(id), nil
}
