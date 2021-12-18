package ports

type DisplayService interface {
	DisplayCustomer(code int) (string, error)
	DisplayInvoice(code string) (string, error)
	DisplayProduct(code string) (string, error)
}