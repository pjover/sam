package ports

type ExternalEditor interface {
	Edit(entity Entity, id string) (string, error)
}

type Entity uint

const (
	Customer Entity = iota
	Invoice
	Product
)
