package invoices

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/viper"
)

type SingleInvoiceGenerator interface {
	Generate(invoiceId string) (string, error)
}

type SingleInvoiceGeneratorImpl struct {
	postManager   hobbit.HttpPostManager
	configService ports.ConfigService
}

func NewSingleInvoiceGenerator() SingleInvoiceGenerator {
	return SingleInvoiceGeneratorImpl{
		hobbit.NewHttpPostManager(),
		cfg.NewConfigService(),
	}
}

func (s SingleInvoiceGeneratorImpl) Generate(invoiceId string) (string, error) {
	fmt.Println("Generant la factura", invoiceId)

	url := fmt.Sprintf("%s/generate/pdf/%s", viper.GetString("urls.hobbit"), invoiceId)

	dirPath, err := s.configService.GetInvoicesDirectory()
	if err != nil {
		return "", err
	}
	return s.postManager.FileWithDefaultName(url, dirPath)
}
