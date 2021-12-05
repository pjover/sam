package env

import "github.com/spf13/viper"

type ConfigManager interface {
	Get(key string) string
	Set(key string, value string) error
}

type configManager struct {
}

func NewConfigManager() ConfigManager {
	return configManager{}
}

func (c configManager) Get(key string) string {
	return viper.GetString(key)
}

func (c configManager) Set(key string, value string) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}
