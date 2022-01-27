package ports

type GenerateService interface {
	CustomerReport() (string, error)
	MonthReport() (string, error)
	ProductReport() (string, error)
}
