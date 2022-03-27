package ports

type CreateService interface {
	CreateCustomer(customerJson []byte) (string, error)
	CreateProduct(productJson []byte) (string, error)
}
