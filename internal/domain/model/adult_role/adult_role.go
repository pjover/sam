package adult_role

type AdultRole uint

const (
	Invalid AdultRole = iota
	Mother
	Father
	Tutor
)

var values = []string{
	"Indefinit",
	"Mare",
	"Pare",
	"Tutor",
}

func (p AdultRole) String() string {
	return values[p]
}
