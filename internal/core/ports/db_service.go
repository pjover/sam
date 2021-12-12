package ports

import "github.com/pjover/sam/internal/core/model"

type DbService interface {
	GetCustomer(code int) (model.Customer, error)
	GetChild(code int) (model.Child, error)
}
