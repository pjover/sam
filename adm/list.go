package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
	"time"
)

type ListManager struct {
	getManager util.HttpGetManager
}

func NewListManager() ListManager {
	return ListManager{
		util.NewHttpGetManager(),
	}
}

func (l ListManager) ListEmails(ei1 bool, ei2 bool, ei3 bool) (string, error) {
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
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListCustomers() (string, error) {
	url := fmt.Sprintf("%s/lists/customers", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListChildren() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListAllCustomersConsumptions() (string, error) {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	url := fmt.Sprintf("%s/consum", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListCustomerConsumptions(customerCode int) (string, error) {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("%s/consum/%d", viper.GetString("urls.hobbit"), customerCode)
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListYearMonthInvoices(yearMonth time.Time) (string, error) {
	ym := yearMonth.Format(util.YearMonthLayout)
	fmt.Println("Llistat de les factures del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByYearMonthIn?yearMonths=%s", viper.GetString("urls.hobbit"), ym)
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListCustomerInvoices(customerCode int) (string, error) {
	fmt.Println("Llistat de les factures del client", customerCode)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerId?customerId=%d", viper.GetString("urls.hobbit"), customerCode)
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) (string, error) {
	ym := yearMonth.Format(util.YearMonthLayout)
	fmt.Println("Llistat de les factures del client", customerCode, "del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerIdAndYearMonthIn?customerId=%d&yearMonths=%s", viper.GetString("urls.hobbit"), customerCode, ym)
	return l.getManager.PrettyJson(url)
}

func (l ListManager) ListProducts() (string, error) {
	fmt.Println("Llistat de tots els productes")
	url := fmt.Sprintf("%s/products?page=0&size=100", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
