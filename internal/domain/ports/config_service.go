package ports

type ConfigService interface {
	Init()
	GetString(key string) string
	SetString(key string, value string) error
	GetWorkingDirectory() (string, error)
	GetInvoicesDirectory() (string, error)
}
