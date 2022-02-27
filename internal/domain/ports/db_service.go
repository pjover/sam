package ports

import "github.com/pjover/sam/internal/domain/model"

type DbService interface {
	FindCustomer(id int) (model.Customer, error)
	FindChild(id int) (model.Child, error)
	FindInvoice(id string) (model.Invoice, error)
	FindProduct(id string) (model.Product, error)
	FindAllProducts() ([]model.Product, error)
	FindInvoicesByYearMonth(yearMonth string) ([]model.Invoice, error)
	FindInvoicesByCustomer(customerId int) ([]model.Invoice, error)
	FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth string) ([]model.Invoice, error)
	FindActiveCustomers() ([]model.Customer, error)
	SearchCustomers(searchText string) ([]model.Customer, error)
	FindActiveChildren() ([]model.Child, error)
	FindAllActiveConsumptions() ([]model.Consumption, error)
	FindActiveChildConsumptions(id int) ([]model.Consumption, error)
	FindAllSequences() ([]model.Sequence, error)
	InsertConsumptions(consumptions []model.Consumption) error
	InsertInvoices(invoices []model.Invoice) error
	UpdateSequences(sequences []model.Sequence) error
	UpdateConsumptions(consumptions []model.Consumption) error
}
