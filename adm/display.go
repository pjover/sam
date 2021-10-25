package adm

import (
	"fmt"
	"sam/comm"
)

func DisplayCustomer(customerCode int) error {
	url := fmt.Sprintf("http://localhost:8080/customers/%d", customerCode)
	return comm.PrintUrl(url)
}

func DisplayProduct(productCode string) error {
	url := fmt.Sprintf("http://localhost:8080/products/%s", productCode)
	return comm.PrintUrl(url)
}

func DisplayInvoice(invoiceCode string) error {
	url := fmt.Sprintf("http://localhost:8080/invoices/%s", invoiceCode)
	return comm.PrintUrl(url)
}
