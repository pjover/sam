package cfg

import (
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/viper"
)

type configService struct {
}

func NewConfigService() ports.ConfigService {
	return configService{}
}

func (c configService) Get(key string) string {
	return viper.GetString(key)
}

func (c configService) Set(key string, value string) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}
