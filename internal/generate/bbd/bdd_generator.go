package bbd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/pjover/sam/internal/generate"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type BddGeneratorImpl struct {
	postManager shared.HttpPostManager
}

func NewBddGenerator() generate.Generator {
	return BddGeneratorImpl{
		shared.NewHttpPostManager(),
	}
}

func (b BddGeneratorImpl) Generate() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)

	dir := shared.GetWorkingDirectory()
	currentFilenames := listFiles(dir, ".qx1")
	filename := getNextBddFilename(currentFilenames)
	return b.postManager.File(url, dir, filename)
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
	for shared.StringInList(filename, currentFilenames) {
		sequence += 1
		filename = buildBddFilename(sequence)
	}
	return filename
}

func buildBddFilename(sequence int) string {
	return fmt.Sprintf("bdd-%d.qx1", sequence)
}
