package generate

type Generator interface {
	Generate() (string, error)
}
