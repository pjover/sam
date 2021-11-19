package generate

import (
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
	"path"
	"path/filepath"
	"sam/internal/util"
)

type GenerateManager interface {
	GenerateBdd() (string, error)
}

type GenerateManagerImpl struct {
	getManager  util.HttpGetManager
	postManager util.HttpPostManager
}

func NewGenerateManager() GenerateManager {
	return GenerateManagerImpl{
		util.NewHttpGetManager(),
		util.NewHttpPostManager(),
	}
}

func (g GenerateManagerImpl) GenerateBdd() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)

	dir := GetWorkingDirectory()
	currentFilenames := listFiles(dir, ".qx1")
	filename := getNextBddFilename(currentFilenames)
	return g.postManager.File(url, dir, filename)
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

func GetWorkingDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}
