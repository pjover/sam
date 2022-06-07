package ports

import "github.com/pjover/sam/internal/domain/model"

type GenerateService interface {
	CustomerReport() (string, error)
	MonthReport(yearMonth model.YearMonth) (string, error)
	ProductReport() (string, error)
	SingleInvoice(id string) (string, error)
	MonthInvoices(yearMonth model.YearMonth) (string, error)
	BddFile() (string, error)
	CustomersCards() (string, error)
}
