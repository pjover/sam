package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core"
	"time"

	"github.com/spf13/viper"
)

type ListInvoices interface {
	ListYearMonthInvoices(yearMonth time.Time) (string, error)
	ListCustomerInvoices(customerCode int) (string, error)
	ListCustomerYearMonthInvoices(customerCode int, yearMonth time.Time) (string, error)
}

type ListInvoicesImpl struct {
	getManager tuk.HttpGetManager
}

func NewListInvoices() ListInvoices {
	return ListInvoicesImpl{
		tuk.NewHttpGetManager(),
	}
}

func (l ListInvoicesImpl) ListYearMonthInvoices(yearMonth time.Time) (string, error) {
	ym := yearMonth.Format(core.YearMonthLayout)
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
	ym := yearMonth.Format(core.YearMonthLayout)
	fmt.Println("Llistat de les factures del client", customerCode, "del mes", ym)
	url := fmt.Sprintf("%s/invoices/search/findByCustomerIdAndYearMonthIn?customerId=%d&yearMonths=%s", viper.GetString("urls.hobbit"), customerCode, ym)
	return l.getManager.PrettyJson(url)
}
