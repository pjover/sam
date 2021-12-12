package ports

type EditService interface {
	EditCustomer(code int) (string, error)
	EditInvoice(code string) (string, error)
	EditProduct(code string) (string, error)
}
