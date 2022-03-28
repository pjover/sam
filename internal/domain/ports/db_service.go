package ports

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"time"
)

type DbService interface {
	FindActiveChildConsumptions(id int) ([]model.Consumption, error)
	FindActiveChildren() ([]model.Child, error)
	FindActiveCustomers() ([]model.Customer, error)
	FindAllActiveConsumptions() ([]model.Consumption, error)
	FindChangedCustomers(changedSince time.Time) ([]model.Customer, error)
	FindAllProducts() ([]model.Product, error)
	FindAllSequences() ([]model.Sequence, error)
	FindChild(id int) (model.Child, error)
	FindCustomer(id int) (model.Customer, error)
	FindInvoice(id string) (model.Invoice, error)
	FindInvoicesByCustomer(customerId int) ([]model.Invoice, error)
	FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth model.YearMonth) ([]model.Invoice, error)
	FindInvoicesByYearMonth(yearMonth model.YearMonth) ([]model.Invoice, error)
	FindInvoicesByYearMonthAndPaymentTypeAndSentToBank(yearMonth model.YearMonth, paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error)
	FindProduct(id string) (model.Product, error)
	InsertConsumptions(consumptions []model.Consumption) error
	InsertInvoices(invoices []model.Invoice) error
	InsertCustomer(customer model.Customer) error
	InsertProduct(product model.Product) error
	UpdateConsumptions(consumptions []model.Consumption) error
	UpdateSequences(sequences []model.Sequence) error
}
