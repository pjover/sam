package lang

import (
	"time"
)

type LangService interface {
	WorkingDir(workingTime time.Time) string
	MonthName(month time.Month) string
}
