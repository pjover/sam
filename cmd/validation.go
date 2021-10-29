package cmd

import (
	"errors"
	"fmt"
	"strconv"
)

func validateCustomerCode(args []string) error {
	err := validateNumberOfArgsEqualsTo(1, args)
	if err != nil {
		return err
	}
	_, err = strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("El codi introduit és invàlid: %s", args[0])
	}
	return nil
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

func validateNumberOfArgsGreaterThan(number int, args []string) error {
	if len(args) > number {
		return fmt.Errorf("Introdueix fins a %d arguments, s'han introduit %d arguments", number, len(args))
	}
	return nil
}

func validateArgsExists(args []string) error {
	if len(args) == 0 {
		return errors.New("Introdueix els arguments, no s'ha introduit cap argument")
	}
	return nil
}
