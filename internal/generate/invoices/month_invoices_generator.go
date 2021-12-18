package invoices

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/core/ports"
	"os"
	"path"

	"github.com/spf13/viper"
)

type MonthInvoicesGenerator interface {
	Generate(onlyNew bool) (string, error)
}

type MonthInvoicesGeneratorImpl struct {
	postManager   hobbit.HttpPostManager
	configService ports.ConfigService
}

func NewMonthInvoicesGenerator() MonthInvoicesGenerator {
	return MonthInvoicesGeneratorImpl{
		hobbit.NewHttpPostManager(),
		cfg.NewConfigService(),
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

	dirPath := path.Join(m.configService.GetWorkingDirectory(), viper.GetString("dirs.invoicesName"))
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	return m.postManager.Zip(url, dirPath)
}
