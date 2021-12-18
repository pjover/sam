package ports

import "github.com/pjover/sam/internal/core/model"

type DbService interface {
	GetCustomer(code int) (model.Customer, error)
	GetChild(code int) (model.Child, error)
	GetInvoice(code string) (model.Invoice, error)
	GetProduct(code string) (model.Product, error)
	GetAllProducts() ([]model.Product, error)
}