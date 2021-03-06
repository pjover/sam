package loader

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
)

type BulkLoader interface {
	LoadMonthInvoices(yearMonth model.YearMonth) (invoices []model.Invoice, err error)
	LoadMonthInvoicesByPaymentType(yearMonth model.YearMonth) (invoices map[payment_type.PaymentType][]model.Invoice, err error)
	LoadProducts() (products map[string]model.Product, err error)
	LoadCustomers() (customers map[int]model.Customer, err error)
	LoadCustomersAndProducts() (customers map[int]model.Customer, products map[string]model.Product, err error)
	LoadMonthInvoicesCustomersAndProducts(yearMonth model.YearMonth) (invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product, err error)
}

type bulkLoader struct {
	configService ports.ConfigService
	dbService     ports.DbService
}

func NewBulkLoader(configService ports.ConfigService, dbService ports.DbService) BulkLoader {
	return bulkLoader{
		configService: configService,
		dbService:     dbService,
	}
}

func (b bulkLoader) LoadMonthInvoices(yearMonth model.YearMonth) (invoices []model.Invoice, err error) {
	invoices, err = b.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return nil, fmt.Errorf("no s'ha pogut carregar les factures des de la base de dades: %s", err)
	}
	return invoices, nil
}

func (b bulkLoader) LoadMonthInvoicesByPaymentType(yearMonth model.YearMonth) (invoices map[payment_type.PaymentType][]model.Invoice, err error) {
	monthInvoices, err := b.LoadMonthInvoices(yearMonth)
	if err != nil {
		return nil, err
	}

	invoices = make(map[payment_type.PaymentType][]model.Invoice)
	for _, invoice := range monthInvoices {
		invoices[invoice.PaymentType()] = append(invoices[invoice.PaymentType()], invoice)
	}
	return invoices, nil
}

func (b bulkLoader) LoadProducts() (products map[string]model.Product, err error) {
	productsList, err := b.dbService.FindAllProducts()
	if err != nil {
		return nil, fmt.Errorf("no s'ha pogut carregar els productes des de la base de dades: %s", err)
	}

	products = make(map[string]model.Product)
	for _, product := range productsList {
		products[product.Id()] = product
	}
	return products, nil
}

func (b bulkLoader) LoadCustomers() (customers map[int]model.Customer, err error) {
	customersList, err := b.dbService.FindActiveCustomers()
	if err != nil {
		return nil, fmt.Errorf("no s'ha pogut carregar els consumidors des de la base de dades: %s", err)
	}

	customers = make(map[int]model.Customer)
	for _, customer := range customersList {
		customers[customer.Id()] = customer
	}
	return customers, nil
}

func (b bulkLoader) LoadCustomersAndProducts() (customers map[int]model.Customer, products map[string]model.Product, err error) {

	customers, err = b.LoadCustomers()
	if err != nil {
		return nil, nil, err
	}

	products, err = b.LoadProducts()
	if err != nil {
		return nil, nil, err
	}
	return customers, products, nil
}

func (b bulkLoader) LoadMonthInvoicesCustomersAndProducts(yearMonth model.YearMonth) (invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product, err error) {

	invoices, err = b.LoadMonthInvoices(yearMonth)
	if err != nil {
		return nil, nil, nil, err
	}

	customers, products, err = b.LoadCustomersAndProducts()
	if err != nil {
		return nil, nil, nil, err
	}
	return invoices, customers, products, nil
}
