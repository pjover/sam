package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/util"
)

type DisplayManager struct {
	getManager util.HttpGetManager
}

func NewDisplayManager() DisplayManager {
	return DisplayManager{
		util.NewHttpGetManager(),
	}
}

func (d DisplayManager) DisplayCustomer(customerCode int) (string, error) {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), customerCode)
	return d.getManager.GetPrint(url)
}

func (d DisplayManager) DisplayProduct(productCode string) (string, error) {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), productCode)
	return d.getManager.GetPrint(url)
}

func (d DisplayManager) DisplayInvoice(invoiceCode string) (string, error) {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), invoiceCode)
	return d.getManager.GetPrint(url)
}
