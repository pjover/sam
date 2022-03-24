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
	FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth string) ([]model.Invoice, error)
	FindInvoicesByYearMonth(yearMonth string) ([]model.Invoice, error)
	FindInvoicesByYearMonthAndPaymentTypeAndSentToBank(yearMonth string, paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error)
	FindProduct(id string) (model.Product, error)
	InsertConsumptions(consumptions []model.Consumption) error
	InsertInvoices(invoices []model.Invoice) error
	SearchCustomers(searchText string) ([]model.Customer, error)
	UpdateConsumptions(consumptions []model.Consumption) error
	UpdateSequences(sequences []model.Sequence) error
}
