package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
	"path"
	"path/filepath"
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

func getNextFilename(filenames []string) string {
	return fmt.Sprintf("bdd-%d.qx1", len(filenames)+1)
}

func find(dir string, ext string) []string {
	var a []string
	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return a
}
