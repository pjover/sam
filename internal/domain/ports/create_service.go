package ports

import "github.com/pjover/sam/internal/domain/model"

type CreateService interface {
	CreateCustomer(customer model.TransientCustomer) (string, error)
	CreateProduct(product model.Product) (string, error)
}
