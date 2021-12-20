package list

import (
	"bytes"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
)

type listService struct {
	dbService ports.DbService
}

func NewListService(dbService ports.DbService) ports.ListService {
	return listService{
		dbService: dbService,
	}
}

func (l listService) ListCustomerInvoices(customerCode int) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomer(customerCode)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func (l listService) ListCustomerYearMonthInvoices(customerCode int, yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomerAndYearMonth(customerCode, yearMonth)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func (l listService) ListYearMonthInvoices(yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func listInvoices(invoices []model.Invoice) (string, error) {
	var buffer bytes.Buffer
	for _, invoice := range invoices {
		buffer.WriteString(fmt.Sprintf("%d  %s  %s  %.2f  %s  %s\n", invoice.CustomerID, invoice.Code, invoice.YearMonth, invoice.Amount(), invoice.PaymentFmt(), invoice.LinesFmt(",")))
	}
	return buffer.String(), nil
}

func (l listService) ListProducts() (string, error) {
	products, err := l.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, product := range products {
		buffer.WriteString(product.String())
	}
	return buffer.String(), nil
}
