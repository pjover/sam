package group_type

type GroupType uint

const (
	Undefined GroupType = iota
	EI_1
	EI_2
	EI_3
)

var values = []string{
	"Indefinit",
	"EI 1 (0-1)",
	"EI 2 (1-2)",
	"EI 3 (2-3)",
}

func (p GroupType) String() string {
	return values[p]
}
