package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
	"path"
	"path/filepath"
	"sam/util"
)

type GenerateManager interface {
	GenerateBdd() (string, error)
	GenerateInvoice(invoiceCode string) (string, error)
	GenerateInvoices() (string, error)
}

type GenerateManagerImpl struct {
	postManager util.HttpPostManager
}

func NewGenerateManager() GenerateManager {
	return GenerateManagerImpl{
		util.NewHttpPostManager(),
	}
}

func (g GenerateManagerImpl) GenerateBdd() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)

	dir := getDirectory()
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

func (g GenerateManagerImpl) GenerateInvoice(invoiceCode string) (string, error) {
	fmt.Println("Generant la factura", invoiceCode)

	url := fmt.Sprintf("%s/generate/pdf/%s", viper.GetString("urls.hobbit"), invoiceCode)

	return g.postManager.FileDefaultName(url, getDirectory())
}

func getDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}

func (g GenerateManagerImpl) GenerateInvoices() (string, error) {
	fmt.Println("Generant les factures del mes")

	url := fmt.Sprintf(
		"%s/generate/pdf?yearMonth=%s&notYetPrinted=false",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)
	return g.postManager.Zip(url, getDirectory(), "factures.zip")
}
