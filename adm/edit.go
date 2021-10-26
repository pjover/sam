package adm

import (
	"fmt"
	"net/url"
	"sam/comm"
)

func EditCustomer(customerCode int) error {
	_url := fmt.Sprintf("http://localhost:8081/db/hobbit_prod/customer/%d", customerCode)
	fmt.Println("Editant el client", customerCode, "a", _url)
	return comm.OpenUrl(_url)
}

func EditProduct(productCode string) error {
	code := url.QueryEscape(fmt.Sprintf("\"%s\"", productCode))
	_url := fmt.Sprintf("http://localhost:8081/db/hobbit_prod/product/%s", code)
	fmt.Println("Editant el producte", productCode, "a", _url)
	return comm.OpenUrl(_url)
}

func EditInvoice(invoiceCode string) error {
	code := url.QueryEscape(fmt.Sprintf("\"%s\"", invoiceCode))
	_url := fmt.Sprintf("http://localhost:8081/db/hobbit_prod/invoice/%s", code)
	fmt.Println("Editant la factura", invoiceCode, "a", _url)
	return comm.OpenUrl(_url)
}
