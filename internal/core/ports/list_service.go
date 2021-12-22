package ports

type ListService interface {
	ListCustomerInvoices(customerCode int) (string, error)
	ListCustomerYearMonthInvoices(customerCode int, yearMonth string) (string, error)
	ListProducts() (string, error)
	ListYearMonthInvoices(yearMonth string) (string, error)
	ListCustomers() (string, error)
	ListChildren() (string, error)
	ListMails() (string, error)
	ListMailsByLanguage() (string, error)
	ListGroupMails(group string) (string, error)
}
