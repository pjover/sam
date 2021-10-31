package adm

import (
	"fmt"
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
	url := fmt.Sprintf("http://localhost:8080/lists/emails/%s", group)
	return comm.PrintGet(url)
}

func ListCustomers() error {
	return comm.PrintGet("http://localhost:8080/lists/customers")
}

func ListChildren() error {
	return comm.PrintGet("http://localhost:8080/lists/children")
}

func ListAllCustomersConsumptions() error {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	return comm.PrintGet("http://localhost:8080/consumptions")
}

func ListCustomerConsumptions(customerCode int) error {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("http://localhost:8080/consumptions/%d", customerCode)
	return comm.PrintGet(url)
}

func ListYearMonthInvoices(yearMonth time.Time) error {
	ym := yearMonth.Format(comm.YearMonthLayout)
	fmt.Println("Llistat de les factures del mes", ym)
	url := fmt.Sprintf("http://localhost:8080/invoices/search/findByYearMonthIn?yearMonths=%s", ym)
	return comm.PrintGet(url)
}

func ListCustomerInvoices(customerCode int) error {
	fmt.Println("Llistat de les factures del client", customerCode)
	url := fmt.Sprintf("http://localhost:8080/invoices/search/findByCustomerId?customerId=%d", customerCode)
	return comm.PrintGet(url)
}

func ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) error {
	ym := yearMonth.Format(comm.YearMonthLayout)
	fmt.Println("Llistat de les factures del client", customerCode, "del mes", ym)
	url := fmt.Sprintf("http://localhost:8080/invoices/search/findByCustomerIdAndYearMonthIn?customerId=%d&yearMonths=%s", customerCode, ym)
	return comm.PrintGet(url)
}

func ListProducts() error {
	fmt.Println("Llistat de tots els productes")
	return comm.PrintGet("http://localhost:8080/products?page=0&size=100")
}
