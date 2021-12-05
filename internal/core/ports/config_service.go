package ports

type ConfigService interface {
	Get(key string) string
	Set(key string, value string) error
}
