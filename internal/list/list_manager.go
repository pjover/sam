package list

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
