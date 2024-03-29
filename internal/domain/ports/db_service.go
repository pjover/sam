package ports

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"time"
)

type DbService interface {
	FindActiveChildConsumptions(id int) ([]model.Consumption, error)
	FindActiveChildren() ([]model.Child, error)
	FindActiveCustomers() ([]model.Customer, error)
	FindAllActiveConsumptions() ([]model.Consumption, error)
	FindAllProducts() ([]model.Product, error)
	FindAllSequences() ([]model.Sequence, error)
	FindChangedCustomers(changedSince time.Time) ([]model.Customer, error)
	FindChild(id int) (model.Child, error)
	FindCustomer(id int) (model.Customer, error)
	FindInvoice(id string) (model.Invoice, error)
	FindInvoicesByCustomer(customerId int) ([]model.Invoice, error)
	FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth model.YearMonth) ([]model.Invoice, error)
	FindInvoicesByYearMonth(yearMonth model.YearMonth) ([]model.Invoice, error)
	FindInvoicesByPaymentTypeAndSentToBank(paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error)
	FindProduct(id string) (model.Product, error)
	FindSequence(sequenceType sequence_type.SequenceType) (model.Sequence, error)
	InsertConsumptions(consumptions []model.Consumption) error
	InsertCustomer(customer model.Customer) error
	InsertInvoices(invoices []model.Invoice) error
	InsertProduct(product model.Product) error
	UpdateConsumptions(consumptions []model.Consumption) error
	UpdateSequences(sequences []model.Sequence) error
	UpdateSequence(sequences model.Sequence) error
	UpdateInvoices(invoices []model.Invoice) error
}
