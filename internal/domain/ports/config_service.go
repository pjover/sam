package ports

import (
	"github.com/pjover/sam/internal/domain/model"
	"time"
)

type ConfigService interface {
	Init()
	GetString(key string) string
	SetString(key string, value string) error
	GetTime(key string) time.Time
	SetTime(key string, value time.Time) error
	GetWorkingDirectory() (string, error)
	GetCurrentYearMonth() model.YearMonth
	SetCurrentYearMonth(yearMonth model.YearMonth) error
	GetInvoicesDirectory() (string, error)
	GetReportsDirectory() (string, error)
	GetCustomersCardsDirectory() (string, error)
}
