package comm

import (
	"fmt"
	"strconv"
)

func ValidateCode(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Introdueix un sol codi, s'han introduit %d arguments", len(args))
	}
	_, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("El codi introduit és invàlid: %s", args[0])
	}
	return nil
}
