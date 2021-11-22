package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
	"time"
)

type ListInvoices interface {
	ListYearMonthInvoices(yearMonth time.Time) (string, error)
	ListCustomerInvoices(customerCode int) (string, error)
	ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) (string, error)
}

type ListInvoicesImpl struct {
	getManager util.HttpGetManager
}

func NewListInvoices() ListInvoices {
	return ListInvoicesImpl{
		util.NewHttpGetManager(),
	}
}

func (l ListInvoicesImpl) ListYearMonthInvoices(yearMonth time.Time) (string, error) {
	ym := yearMonth.Format(util.YearMonthLayout)
	fmt.Println("Llistat de les factures del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByYearMonthIn?yearMonths=%s", viper.GetString("urls.hobbit"), ym)
	return l.getManager.PrettyJson(url)
}

func (l ListInvoicesImpl) ListCustomerInvoices(customerCode int) (string, error) {
	fmt.Println("Llistat de les factures del client", customerCode)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerId?customerId=%d", viper.GetString("urls.hobbit"), customerCode)
	return l.getManager.PrettyJson(url)
}

func (l ListInvoicesImpl) ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) (string, error) {
	ym := yearMonth.Format(util.YearMonthLayout)
	fmt.Println("Llistat de les factures del client", customerCode, "del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerIdAndYearMonthIn?customerId=%d&yearMonths=%s", viper.GetString("urls.hobbit"), customerCode, ym)
	return l.getManager.PrettyJson(url)
}
