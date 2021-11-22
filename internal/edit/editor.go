package edit

type Editor interface {
	Edit(code string) error
}
