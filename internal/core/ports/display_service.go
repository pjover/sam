package ports

type DisplayService interface {
	DisplayCustomer(code int) (string, error)
}
