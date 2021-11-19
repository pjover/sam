package invoices

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type SingleInvoiceGenerator interface {
	Generate(invoiceCode string) (string, error)
}

type SingleInvoiceGeneratorImpl struct {
	postManager util.HttpPostManager
}

func NewSingleInvoiceGenerator() SingleInvoiceGenerator {
	return SingleInvoiceGeneratorImpl{
		util.NewHttpPostManager(),
	}
}

func (s SingleInvoiceGeneratorImpl) Generate(invoiceCode string) (string, error) {
	fmt.Println("Generant la factura", invoiceCode)

	url := fmt.Sprintf("%s/generate/pdf/%s", viper.GetString("urls.hobbit"), invoiceCode)

	return s.postManager.FileWithDefaultName(url, util.GetWorkingDirectory())
}
