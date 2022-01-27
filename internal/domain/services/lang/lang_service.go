package lang

import (
	"time"
)

type LangService interface {
	WorkingDir(month time.Time) string
	MonthName(month time.Time) string
}
