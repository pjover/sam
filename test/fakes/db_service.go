package fakes

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"github.com/pjover/sam/internal/domain/ports"
	"time"
)

type dbService struct {
	customers    map[int]model.Customer
	children     map[int]model.Child
	products     map[string]model.Product
	sequences    map[sequence_type.SequenceType]model.Sequence
	consumptions map[string]model.Consumption
	invoices     map[string]model.Invoice
}

func FakeDbService() ports.DbService {
	customers := loadCustomers()
	return &dbService{
		customers:    customers,
		children:     loadChildren(customers),
		products:     loadProducts(),
		sequences:    loadSequences(),
		consumptions: loadConsumptions(),
		invoices:     loadInvoices(),
	}
}

func loadCustomers() map[int]model.Customer {
	var customers = make(map[int]model.Customer)
	customers[148] = model.TestCustomer148
	customers[149] = model.TestCustomer149
	return customers
}

func loadChildren(customers map[int]model.Customer) map[int]model.Child {
	var children = make(map[int]model.Child)
	for _, customer := range customers {
		for _, child := range customer.Children() {
			children[child.Id()] = child
		}
	}
	return children
}

func loadProducts() map[string]model.Product {
	var products = make(map[string]model.Product)
	products["TST"] = model.ProductTST
	products["XXX"] = model.ProductXXX
	products["YYY"] = model.ProductYYY
	return products
}

func loadSequences() map[sequence_type.SequenceType]model.Sequence {
	var sequences = make(map[sequence_type.SequenceType]model.Sequence)
	sequences[sequence_type.Customer] = model.NewSequence(sequence_type.Customer, 150)
	sequences[sequence_type.RectificationInvoice] = model.NewSequence(sequence_type.RectificationInvoice, 11)
	sequences[sequence_type.SpecialInvoice] = model.NewSequence(sequence_type.SpecialInvoice, 22)
	sequences[sequence_type.StandardInvoice] = model.NewSequence(sequence_type.StandardInvoice, 33)
	return sequences
}

func loadConsumptions() map[string]model.Consumption {
	return make(map[string]model.Consumption)
}

func loadInvoices() map[string]model.Invoice {
	return make(map[string]model.Invoice)
}

func (d dbService) FindActiveChildConsumptions(id int) ([]model.Consumption, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindActiveChildren() ([]model.Child, error) {
	var children []model.Child
	for _, child := range d.children {
		children = append(children, child)
	}
	return children, nil
}

func (d dbService) FindActiveCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	for _, customer := range d.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}

func (d dbService) FindAllActiveConsumptions() ([]model.Consumption, error) {
	var consumptions []model.Consumption
	for _, consumption := range d.consumptions {
		if consumption.InvoiceId() == "" {
			continue
		}
		consumptions = append(consumptions, consumption)
	}
	return consumptions, nil
}

func (d dbService) FindAllProducts() ([]model.Product, error) {
	var products []model.Product
	for _, product := range d.products {
		products = append(products, product)
	}
	return products, nil
}

func (d dbService) FindAllSequences() ([]model.Sequence, error) {
	var sequences []model.Sequence
	for _, sequence := range d.sequences {
		sequences = append(sequences, sequence)
	}
	return sequences, nil
}

func (d dbService) FindChangedCustomers(changedSince time.Time) ([]model.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindChild(id int) (model.Child, error) {
	child, exists := d.children[id]
	if !exists {
		return model.Child{}, fmt.Errorf("no s'ha trobat l'infant amb codi %d", id)
	}
	return child, nil
}

func (d dbService) FindCustomer(id int) (model.Customer, error) {
	customer, exists := d.customers[id]
	if !exists {
		return model.Customer{}, fmt.Errorf("no s'ha trobat el client amb codi %d", id)
	}
	return customer, nil
}

func (d dbService) FindInvoice(id string) (model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindInvoicesByCustomer(customerId int) ([]model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth model.YearMonth) ([]model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindInvoicesByYearMonth(yearMonth model.YearMonth) ([]model.Invoice, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbService) FindInvoicesByPaymentTypeAndSentToBank(paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error) {
	var invoices []model.Invoice
	for _, invoice := range d.invoices {
		if invoice.PaymentType() == paymentType && invoice.SentToBank() == sentToBank {
			invoices = append(invoices, invoice)
		}
	}
	return invoices, nil
}

func (d dbService) FindProduct(id string) (model.Product, error) {
	product, exists := d.products[id]
	if !exists {
		return model.Product{}, fmt.Errorf("no s'ha trobat el producte amb codi %s", id)
	}
	return product, nil
}

func (d dbService) FindSequence(sequenceType sequence_type.SequenceType) (model.Sequence, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbService) InsertConsumptions(consumptions []model.Consumption) error {
	for _, consumption := range consumptions {
		d.consumptions[consumption.Id()] = consumption
	}
	return nil
}

func (d dbService) InsertCustomer(customer model.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (d *dbService) InsertInvoices(invoices []model.Invoice) error {
	for _, invoice := range invoices {
		d.invoices[invoice.Id()] = invoice
	}
	return nil
}

func (d dbService) InsertProduct(product model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (d dbService) UpdateConsumptions(consumptions []model.Consumption) error {
	for _, consumption := range consumptions {
		d.consumptions[consumption.Id()] = consumption
	}
	return nil
}

func (d dbService) UpdateSequences(sequences []model.Sequence) error {
	for _, sequence := range sequences {
		d.sequences[sequence.Id()] = sequence
	}
	return nil
}

func (d dbService) UpdateSequence(sequences model.Sequence) error {
	//TODO implement me
	panic("implement me")
}

func (d dbService) UpdateInvoices(invoices []model.Invoice) error {
	for _, invoice := range invoices {
		d.invoices[invoice.Id()] = invoice
	}
	return nil
}
