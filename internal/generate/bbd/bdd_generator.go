package bbd

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/generate"
	"io/fs"
	"path/filepath"
)

type BddGeneratorImpl struct {
	configService ports.ConfigService
	postManager   hobbit.HttpPostManager
}

func NewBddGenerator(configService ports.ConfigService) generate.Generator {
	return BddGeneratorImpl{
		configService: configService,
		postManager:   hobbit.NewHttpPostManager(),
	}
}

func (b BddGeneratorImpl) Generate() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		b.configService.GetString("urls.hobbit"),
		b.configService.GetString("yearMonth"),
	)

	dir, err := b.configService.GetWorkingDirectory()
	if err != nil {
		return "", err
	}
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
	for stringInList(filename, currentFilenames) {
		sequence += 1
		filename = buildBddFilename(sequence)
	}
	return filename
}

func stringInList(str string, list []string) bool {
	for _, b := range list {
		if b == str {
			return true
		}
	}
	return false
}

func buildBddFilename(sequence int) string {
	return fmt.Sprintf("bdd-%d.qx1", sequence)
}
