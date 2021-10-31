package cmd

import (
	"errors"
	"fmt"
	"sam/comm"
	"strconv"
	"time"
)

func validateCustomerCode(args []string) error {
	err := validateNumberOfArgsEqualsTo(1, args)
	if err != nil {
		return err
	}
	_, err = parseInteger(args[0], "de client")
	return err
}

func validateProductCode(args []string) error {
	err := validateNumberOfArgsEqualsTo(1, args)
	if err != nil {
		return err
	}
	if len(args[0]) != 3 {
		return fmt.Errorf("El codi introduit és invàlid: %s", args[0])
	}
	return nil
}

func validateNumberOfArgsEqualsTo(number int, args []string) error {
	if len(args) != number {
		return fmt.Errorf("Introdueix %d arguments, s'han introduit %d arguments", number, len(args))
	}
	return nil
}

func validateNumberOfArgsGreaterThan(min int, args []string) error {
	if len(args) < min {
		return fmt.Errorf("Introdueix més de %d arguments, has introduit %d arguments", min, len(args))
	}
	return nil
}
func validateNumberOfArgsBetween(min int, max int, args []string) error {
	if len(args) < min && len(args) > max {
		return fmt.Errorf("Introdueix des de %d fins a %d arguments, has introduit %d arguments", min, max, len(args))
	}
	return nil
}

func validateArgsExists(args []string) error {
	if len(args) == 0 {
		return errors.New("Introdueix els arguments, no s'ha introduit cap argument")
	}
	return nil
}

func parseInteger(customCode string, codeType string) (int, error) {
	code, err := strconv.Atoi(customCode)
	if err != nil {
		return 0, fmt.Errorf("El codi %s introduit és invàlid: %s", codeType, customCode)
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
	ym, err := time.Parse(comm.YearMonthLayout, yearMonth)
	if err != nil {
		return time.Time{}, errors.New("Error al introduïr el mes: " + err.Error())
	}
	return ym, nil
}
