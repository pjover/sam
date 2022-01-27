package ports

type SearchService interface {
	SearchCustomer(searchText string) (string, error)
}
