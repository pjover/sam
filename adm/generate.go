package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"sam/util"
)

type GenerateManager struct {
	postManager util.HttpPostManager
}

func NewGenerateManager() GenerateManager {
	return GenerateManager{
		util.NewHttpPostManager(),
	}
}

func (g GenerateManager) GenerateBdd() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)
	//TODO create file sequence
	filePath := path.Join(
		viper.GetString("dirs.home"),
		viper.GetString("dirs.current"),
		"bdd.q1x",
	)
	return g.postManager.File(url, filePath)
}
