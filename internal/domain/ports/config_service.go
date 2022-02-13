package ports

type ConfigService interface {
	Init()
	Get(key string) string
	Set(key string, value string) error
	GetWorkingDirectory() (string, error)
	GetInvoicesDirectory() (string, error)
}