package list

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/loader"
	"sort"
)

type listService struct {
	configService ports.ConfigService
	dbService     ports.DbService
	bulkLoader    loader.BulkLoader
}

func NewListService(configService ports.ConfigService, dbService ports.DbService, bulkLoader loader.BulkLoader) ports.ListService {
	return listService{
		configService: configService,
		dbService:     dbService,
		bulkLoader:    bulkLoader,
	}
}

func (l listService) ListCustomerInvoices(customerId int) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomer(customerId)
	if err != nil {
		return "", err
	}
	titleMessage := fmt.Sprintf("Lists of customer %d invoices:", customerId)
	return listInvoices(titleMessage, invoices)
}

func (l listService) ListCustomerYearMonthInvoices(customerId int, yearMonth model.YearMonth) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomerAndYearMonth(customerId, yearMonth)
	if err != nil {
		return "", err
	}
	titleMessage := fmt.Sprintf("Lists of customer %d and %s year-month invoices:", customerId, yearMonth)
	return listInvoices(titleMessage, invoices)
}

func (l listService) ListYearMonthInvoices(yearMonth model.YearMonth) (string, error) {
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
		buffer.WriteString(customer.InvoiceHolder().Mail() + ", ")
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
		if customer.Language() == language.Catalan {
			caBuffer.WriteString(customer.InvoiceHolder().Mail() + ", ")
		} else {
			esBuffer.WriteString(customer.InvoiceHolder().Mail() + ", ")
		}
	}
	return caBuffer.String() + "\n" + esBuffer.String(), nil
}

func (l listService) ListGroupMails(groupType group_type.GroupType) (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(groupType.Format() + ":\n")
	for _, customer := range customers {
		var in bool
		for _, child := range customer.Children() {
			if child.Group() == groupType {
				in = true
				break
			}
		}
		if in {
			buffer.WriteString(customer.InvoiceHolder().Mail() + ", ")
		}
	}
	return buffer.String(), nil
}

func (l listService) ListConsumptions() (string, error) {
	consumptions, err := l.dbService.FindAllActiveConsumptions()
	if err != nil {
		return "", err
	}

	children, err := l.dbService.FindActiveChildren()
	if err != nil {
		return "", err
	}

	products, err := l.bulkLoader.LoadProducts()
	if err != nil {
		return "", err
	}

	customers, err := l.loadSortedCustomers()
	if err != nil {
		return "", err
	}

	return l.getConsumptionsText(consumptions, children, products, customers), nil
}

func (l listService) loadSortedCustomers() ([]model.Customer, error) {
	customers, err := l.bulkLoader.LoadCustomers()
	if err != nil {
		return nil, err
	}

	var customersList []model.Customer
	for _, customer := range customers {
		customersList = append(customersList, customer)
	}

	sort.Slice(customersList, func(i, j int) bool {
		return customersList[i].Id() < customersList[j].Id()
	})

	return customersList, nil
}

func (l listService) getConsumptionsText(consumptions []model.Consumption, children []model.Child, products map[string]model.Product, customers []model.Customer) string {
	var buffer bytes.Buffer

	var totalAmount float64
	for paymentType := payment_type.Invalid; paymentType <= payment_type.Rectification; paymentType++ {
		var childCounter int
		var parcialAmount float64
		for _, customer := range customers {
			if customer.InvoiceHolder().PaymentType() != paymentType {
				continue
			}
			for _, child := range children {
				if child.CustomerId() != customer.Id() {
					continue
				}
				text, total := l.getChildValues(child, consumptions, products)
				if text == "" {
					continue
				}
				buffer.WriteString(text)
				childCounter += 1
				parcialAmount += total
				totalAmount += total
			}
		}
		if parcialAmount != 0.0 {
			buffer.WriteString(fmt.Sprintf("\n%d infant(s) amb %s: %.02f €\n\n", childCounter, paymentType.Format(), parcialAmount))
		}
	}
	buffer.WriteString(fmt.Sprintf("TOTAL: %.02f €", totalAmount))
	return buffer.String()
}

func (l listService) getChildValues(child model.Child, consumptions []model.Consumption, products map[string]model.Product) (string, float64) {
	var cons []model.Consumption
	for _, c := range consumptions {
		if c.ChildId() == child.Id() {
			cons = append(cons, c)
		}
	}
	if len(cons) > 0 {
		return model.ConsumptionListFormatValues(consumptions, child, products, "  ")
	}
	return "", 0
}

func (l listService) ListChildConsumptions(childId int) (string, error) {
	consumptions, err := l.dbService.FindActiveChildConsumptions(childId)
	if err != nil {
		return "", err
	}

	child, err := l.dbService.FindChild(childId)
	if err != nil {
		return "", err
	}

	products, err := l.bulkLoader.LoadProducts()
	if err != nil {
		return "", err
	}

	text, _ := model.ConsumptionListFormatValues(consumptions, child, products, "")
	return text, nil
}
