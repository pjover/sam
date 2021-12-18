package ports

type GenerateService interface {
	GenerateProduct() (string, error)
}
