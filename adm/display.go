package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/comm"
)

func DisplayCustomer(customerCode int) error {
	url := fmt.Sprintf("%s/customers/%d", viper.GetString("urls.hobbit"), customerCode)
	return comm.PrintGet(url)
}

func DisplayProduct(productCode string) error {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), productCode)
	return comm.PrintGet(url)
}

func DisplayInvoice(invoiceCode string) error {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), invoiceCode)
	return comm.PrintGet(url)
}
