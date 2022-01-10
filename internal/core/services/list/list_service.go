package list

import (
	"bytes"
	"fmt"
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
	titleMessage := fmt.Sprintf("Lists of customer %d invoices:", customerCode)
	return listInvoices(titleMessage, invoices)
}

func (l listService) ListCustomerYearMonthInvoices(customerCode int, yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomerAndYearMonth(customerCode, yearMonth)
	if err != nil {
		return "", err
	}
	titleMessage := fmt.Sprintf("Lists of customer %d and %s year-month invoices:", customerCode, yearMonth)
	return listInvoices(titleMessage, invoices)
}

func (l listService) ListYearMonthInvoices(yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return "", err
	}
	titleMessage := fmt.Sprintf("Lists of all %s year-month invoices:", yearMonth)
	return listInvoices(titleMessage, invoices)
}

func listInvoices(titleMessage string, invoices []model.Invoice) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(titleMessage + "\n")
	for _, invoice := range invoices {
		buffer.WriteString(" - " + invoice.String() + "\n")
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
		buffer.WriteString(product.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListCustomers() (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, customer := range customers {
		buffer.WriteString(customer.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListChildren() (string, error) {
	children, err := l.dbService.FindActiveChildren()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, child := range children {
		buffer.WriteString(child.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListMails() (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, customer := range customers {
		buffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
	}
	return buffer.String(), nil
}

func (l listService) ListMailsByLanguage() (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var caBuffer, esBuffer bytes.Buffer
	caBuffer.WriteString("CA:\n")
	esBuffer.WriteString("ES:\n")
	for _, customer := range customers {
		if customer.Language == "CA" {
			caBuffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		} else {
			esBuffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		}
	}
	return caBuffer.String() + "\n" + esBuffer.String(), nil
}

func (l listService) ListGroupMails(group string) (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(group + ":\n")
	for _, customer := range customers {
		var in bool
		for _, child := range customer.Children {
			if child.Group == group {
				in = true
				break
			}
		}
		if in {
			buffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		}
	}
	return buffer.String(), nil
}

func (l listService) ListConsumptions() (string, error) {
	consumptions, err := l.dbService.FindAllConsumptions()
	if err != nil {
		return "", err
	}
	return l.printConsumptions(consumptions)
}

func (l listService) ListChildConsumptions(childCode int) (string, error) {
	consumptions, err := l.dbService.FindChildConsumptions(childCode)
	if err != nil {
		return "", err
	}
	return l.printConsumptions(consumptions)
}

func (l listService) printConsumptions(consumptions []model.Consumption) (string, error) {
	var buffer bytes.Buffer
	for _, consumption := range consumptions {
		buffer.WriteString(consumption.String() + "\n")
	}
	return buffer.String(), nil
}
