package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"sam/util"
	"strconv"
	"strings"
	"time"
)

func validateNumberOfArgsEqualsTo(number int, args []string) error {
	if len(args) != number {
		return fmt.Errorf("Introdueix %d arguments, s'han introduit %d arguments", number, len(args))
	}
	return nil
}

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

func parseInteger(strCode string, codeType string) (int, error) {
	code, err := strconv.Atoi(strCode)
	if err != nil {
		return 0, fmt.Errorf("El codi %s introduit és invàlid: %s", codeType, strCode)
	}
	return code, nil
}

func parseFloat(value string) (float64, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("El número introduit és invàlid: %s", value)
	}
	return float, nil
}

func parseYearMonth(yearMonth string) (time.Time, error) {
	ym, err := time.Parse(util.YearMonthLayout, yearMonth)
	if err != nil {
		return time.Time{}, errors.New("Error al introduïr el mes: " + err.Error())
	}
	return ym, nil
}

func parseProductCode(code string) (string, error) {
	if len(code) != 3 {
		return "", fmt.Errorf("El codi de producte introduit és invàlid: %s", code)
	}
	return strings.ToUpper(code), nil
}
