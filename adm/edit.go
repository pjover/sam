package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/internal/util"
)

func EditCustomer(customerCode int) error {

	_url := fmt.Sprintf("%s/customer/%d", viper.GetString("urls.mongoExpress"), customerCode)
	fmt.Println("Editant el client", customerCode, "a", _url)
	return util.OpenOnBrowser(_url)
}

func EditProduct(productCode string) error {
	code := url.QueryEscape(fmt.Sprintf("\"%s\"", productCode))
	_url := fmt.Sprintf("%s/product/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el producte", productCode, "a", _url)
	return util.OpenOnBrowser(_url)
}

func EditInvoice(invoiceCode string) error {
	code := url.QueryEscape(fmt.Sprintf("\"%s\"", invoiceCode))
	_url := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant la factura", invoiceCode, "a", _url)
	return util.OpenOnBrowser(_url)
}
