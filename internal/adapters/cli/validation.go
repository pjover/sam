package cli

import (
	"fmt"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
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

func ParseInteger(strCode string, codeType string) (int, error) {
	code, err := strconv.Atoi(strCode)
	if err != nil {
		return 0, fmt.Errorf("el codi %s introduit és invàlid: %s", codeType, strCode)
	}
	return code, nil
}

func ParseFloat(value string) (float64, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("el número introduit és invàlid: %s", value)
	}
	return float, nil
}

func ParseYearMonth(yearMonth string) (time.Time, error) {
	ym, err := time.Parse(shared.YearMonthLayout, yearMonth)
	if err != nil {
		return time.Time{}, fmt.Errorf("error al introduïr el mes: %s", err.Error())
	}
	return ym, nil
}

func ParseProductCode(code string) (string, error) {
	if len(code) != 3 {
		return "", fmt.Errorf("el codi de producte introduit és invàlid: %s", code)
	}
	return strings.ToUpper(code), nil
}
