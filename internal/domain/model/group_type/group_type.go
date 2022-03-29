package group_type

import "strings"

type GroupType uint

const (
	Undefined GroupType = iota
	EI_1
	EI_2
	EI_3
)

var stringValues = []string{
	"UNDEFINED",
	"EI_1",
	"EI_2",
	"EI_3",
}

func (g GroupType) String() string {
	return stringValues[g]
}

func NewGroupType(value string) GroupType {

	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return GroupType(i)
		}
	}
	return Undefined
}

var formatValues = []string{
	"Indefinit",
	"EI 1 (0-1)",
	"EI 2 (1-2)",
	"EI 3 (2-3)",
}

func (g GroupType) Format() string {
	return formatValues[g]
}
