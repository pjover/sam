package shared

import (
	"github.com/spf13/viper"
	"path"
)

// TODO Remove

func GetWorkingDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}
