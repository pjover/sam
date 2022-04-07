package adult_role

import "strings"

type AdultRole uint

const (
	Invalid AdultRole = iota
	Mother
	Father
	Tutor
)

var stringValues = []string{
	"",
	"MOTHER",
	"FATHER",
	"TUTOR",
}

func (p AdultRole) String() string {
	return stringValues[p]
}

func NewAdultRole(value string) AdultRole {
	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return AdultRole(i)
		}
	}
	return Invalid
}

var formatValues = []string{
	"Indefinit",
	"Mare",
	"Pare",
	"Tutor",
}

func (p AdultRole) Format() string {
	return formatValues[p]
}
