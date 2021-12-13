package ports

type ExternalEditor interface {
	Edit(entity Entity, code string) (string, error)
}

type Entity uint

const (
	Customer Entity = iota
	Invoice
	Product
)
