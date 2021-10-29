package adm

import (
	"errors"
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
	return comm.PrintUrl(url)
}

func ListCustomers() error {
	return comm.PrintUrl("http://localhost:8080/lists/customers")
}

func ListChildren() error {
	return comm.PrintUrl("http://localhost:8080/lists/children")
}

func ListAllCustomersConsumptions() error {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	return comm.PrintUrl("http://localhost:8080/consumptions")
}

func ListCustomerConsumptions(customerCode int) error {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("http://localhost:8080/consumptions/%d", customerCode)
	return comm.PrintUrl(url)
}

func ListYearMonthInvoices(yearMonth string) error {
	const layout = "2006-01"
	if yearMonth == "" {
		yearMonth = time.Now().Format(layout)
	} else {
		_, err := time.Parse(layout, yearMonth)
		if err != nil {
			return errors.New("Error al introduir el mes: " + err.Error())
		}
	}
	fmt.Println("Llistat de les factures del mes", yearMonth)
	url := fmt.Sprintf("http://localhost:8080/invoices/search/findByYearMonthIn?yearMonths=%s", yearMonth)
	return comm.PrintUrl(url)
}
