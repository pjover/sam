package create

import "github.com/pjover/sam/internal/domain/ports"

type createService struct {
	dbService ports.DbService
}

func NewCreateService(dbService ports.DbService) ports.CreateService {
	return createService{
		dbService: dbService,
	}
}

func (c createService) CreateCustomer(customerJson []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c createService) CreateProduct(productJson []byte) (string, error) {
	//TODO implement me
	panic("implement me")
}
