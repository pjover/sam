package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/util"
)

func DisplayCustomer(customerCode int) error {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), customerCode)
	return util.PrintGet(url)
}

func DisplayProduct(productCode string) error {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), productCode)
	return util.PrintGet(url)
}

func DisplayInvoice(invoiceCode string) error {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), invoiceCode)
	return util.PrintGet(url)
}
