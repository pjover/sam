package ports

type DisplayService interface {
	DisplayCustomer(id int) (string, error)
	DisplayInvoice(id string) (string, error)
	DisplayProduct(id string) (string, error)
}
