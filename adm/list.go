package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/comm"
	"time"
)

func ListEmails(ei1 bool, ei2 bool, ei3 bool) error {
	var group string
	if ei1 {
		group = "EI_1"
	} else if ei2 {
		group = "EI_2"
	} else if ei3 {
		group = "EI_3"
	} else {
		group = "ALL"
	}

	url := fmt.Sprintf("%s/lists/emails/%s", viper.GetString("urls.hobbit"), group)
	return comm.PrintGet(url)
}

func ListCustomers() error {
	url := fmt.Sprintf("%s/lists/customers", viper.GetString("urls.hobbit"))
	return comm.PrintGet(url)
}

func ListChildren() error {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return comm.PrintGet(url)
}

func ListAllCustomersConsumptions() error {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	url := fmt.Sprintf("%s/consumptions", viper.GetString("urls.hobbit"))
	return comm.PrintGet(url)
}

func ListCustomerConsumptions(customerCode int) error {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("%s/consumptions/%d", viper.GetString("urls.hobbit"), customerCode)
	return comm.PrintGet(url)
}

func ListYearMonthInvoices(yearMonth time.Time) error {
	ym := yearMonth.Format(comm.YearMonthLayout)
	fmt.Println("Llistat de les factures del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByYearMonthIn?yearMonths=%s", viper.GetString("urls.hobbit"), ym)
	return comm.PrintGet(url)
}

func ListCustomerInvoices(customerCode int) error {
	fmt.Println("Llistat de les factures del client", customerCode)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerId?customerId=%d", viper.GetString("urls.hobbit"), customerCode)
	return comm.PrintGet(url)
}

func ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) error {
	ym := yearMonth.Format(comm.YearMonthLayout)
	fmt.Println("Llistat de les factures del client", customerCode, "del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerIdAndYearMonthIn?customerId=%d&yearMonths=%s", viper.GetString("urls.hobbit"), customerCode, ym)
	return comm.PrintGet(url)
}

func ListProducts() error {
	fmt.Println("Llistat de tots els productes")
	url := fmt.Sprintf("%s/products?page=0&size=100", viper.GetString("urls.hobbit"))
	return comm.PrintGet(url)
}
