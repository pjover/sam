package ports

import "github.com/pjover/sam/internal/domain/model/group_type"

type ListService interface {
	ListCustomerInvoices(customerId int) (string, error)
	ListCustomerYearMonthInvoices(customerId int, yearMonth string) (string, error)
	ListProducts() (string, error)
	ListYearMonthInvoices(yearMonth string) (string, error)
	ListCustomers() (string, error)
	ListChildren() (string, error)
	ListMails() (string, error)
	ListMailsByLanguage() (string, error)
	ListGroupMails(groupType group_type.GroupType) (string, error)
	ListConsumptions() (string, error)
	ListChildConsumptions(childId int) (string, error)
}
