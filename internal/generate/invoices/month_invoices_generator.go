package invoices

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sam/internal/util"
)

type MonthInvoicesGenerator interface {
	Generate(onlyNew bool) (string, error)
}

type MonthInvoicesGeneratorImpl struct {
	postManager util.HttpPostManager
}

func NewMonthInvoicesGenerator() MonthInvoicesGenerator {
	return MonthInvoicesGeneratorImpl{
		util.NewHttpPostManager(),
	}
}

func (m MonthInvoicesGeneratorImpl) Generate(onlyNew bool) (string, error) {
	fmt.Println("Generant les factures del mes")

	url := fmt.Sprintf(
		"%s/generate/pdf?yearMonth=%s&notYetPrinted=%t",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
		onlyNew,
	)

	dirPath := path.Join(util.GetWorkingDirectory(), viper.GetString("dirs.invoicesName"))
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	return m.postManager.Zip(url, dirPath)
}
