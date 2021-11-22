package display

type Display interface {
	Display(code string) (string, error)
}
