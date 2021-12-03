package invoices

import (
	"fmt"
	"os"
	"path"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type MonthInvoicesGenerator interface {
	Generate(onlyNew bool) (string, error)
}

type MonthInvoicesGeneratorImpl struct {
	postManager shared.HttpPostManager
}

func NewMonthInvoicesGenerator() MonthInvoicesGenerator {
	return MonthInvoicesGeneratorImpl{
		shared.NewHttpPostManager(),
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

	dirPath := path.Join(shared.GetWorkingDirectory(), viper.GetString("dirs.invoicesName"))
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	return m.postManager.Zip(url, dirPath)
}
