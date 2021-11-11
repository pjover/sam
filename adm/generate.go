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

	dir := path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
	currentFilenames := listFiles(dir, ".qx1")
	filename := getNextBddFilename(currentFilenames)
	filePath := path.Join(dir, filename)
	return g.postManager.File(url, filePath)
}

func listFiles(dir string, ext string) []string {
	var filenames []string
	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == ext {
			filenames = append(filenames, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return filenames
}

func getNextBddFilename(currentFilenames []string) string {
	sequence := len(currentFilenames) + 1
	filename := buildBddFilename(sequence)
	for util.StringInList(filename, currentFilenames) {
		sequence += 1
		filename = buildBddFilename(sequence)
	}
	return filename
}

func buildBddFilename(sequence int) string {
	return fmt.Sprintf("bdd-%d.qx1", sequence)
}
