package search

import (
	"bytes"
	"github.com/pjover/sam/internal/core/ports"
)

type searchService struct {
	dbService ports.DbService
}

func NewSearchService(dbService ports.DbService) ports.SearchService {
	return searchService{
		dbService: dbService,
	}
}

func (s searchService) SearchCustomer(searchText string) (string, error) {
	customers, err := s.dbService.SearchCustomers(searchText)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, customer := range customers {
		buffer.WriteString(customer.String() + "\n")
	}
	return buffer.String(), nil
}
