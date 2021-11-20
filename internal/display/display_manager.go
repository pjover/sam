package display

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type DisplayManager struct {
	getManager util.HttpGetManager
}

func NewDisplayManager() DisplayManager {
	return DisplayManager{
		util.NewHttpGetManager(),
	}
}

func (d DisplayManager) DisplayProduct(productCode string) (string, error) {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), productCode)
	return d.getManager.PrettyJson(url)
}

func (d DisplayManager) DisplayInvoice(invoiceCode string) (string, error) {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), invoiceCode)
	return d.getManager.PrettyJson(url)
}
