package list

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/core/ports"
)

type listService struct {
	dbService ports.DbService
}

func NewListService(dbService ports.DbService) ports.ListService {
	return listService{
		dbService: dbService,
	}
}

func (l listService) ListProducts() (string, error) {
	products, err := l.dbService.GetAllProducts()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, product := range products {
		buffer.WriteString(fmt.Sprintf("%s  %.2f \t %s\n", product.Id, product.Price, product.Name))
	}
	return buffer.String(), nil
}
