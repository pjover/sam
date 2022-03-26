package bdd

type stringBddBuilder struct {
}

func NewStringBddBuilder() BddBuilder {
	return stringBddBuilder{}
}

func (s stringBddBuilder) Build(bdd Bdd) (content string, err error) {
	//TODO implement me
	panic("implement me")
}
