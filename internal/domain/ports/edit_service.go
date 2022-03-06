package ports

type EditService interface {
	EditCustomer(id int) (string, error)
	EditInvoice(id string) (string, error)
	EditProduct(id string) (string, error)
}
