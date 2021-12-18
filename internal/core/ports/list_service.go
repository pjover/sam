package ports

type ListService interface {
	ListProducts() (string, error)
}
