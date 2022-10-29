package ports

import "github.com/pjover/sam/internal/domain/model"

type BillingService interface {
	InsertConsumptions(id int, consumptions map[string]float64, yearMonth model.YearMonth, note string) (string, error)
	BillConsumptions() (string, error)
	RectifyConsumptions(id int, consumptions map[string]float64, yearMonth model.YearMonth, note string) (string, error)
}
