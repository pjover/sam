package invoices

import (
	"fmt"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type SingleInvoiceGenerator interface {
	Generate(invoiceCode string) (string, error)
}

type SingleInvoiceGeneratorImpl struct {
	postManager shared.HttpPostManager
}

func NewSingleInvoiceGenerator() SingleInvoiceGenerator {
	return SingleInvoiceGeneratorImpl{
		shared.NewHttpPostManager(),
	}
}

func (s SingleInvoiceGeneratorImpl) Generate(invoiceCode string) (string, error) {
	fmt.Println("Generant la factura", invoiceCode)

	url := fmt.Sprintf("%s/generate/pdf/%s", viper.GetString("urls.hobbit"), invoiceCode)

	return s.postManager.FileWithDefaultName(url, shared.GetWorkingDirectory())
}
