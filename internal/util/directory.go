package util

import (
	"github.com/spf13/viper"
	"path"
)

func GetWorkingDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}
