package bdd

type BddBuilder interface {
	Build(bdd Bdd) (content string, err error)
}
