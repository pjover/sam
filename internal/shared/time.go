package shared

import "time"

type TimeManager interface {
	Now() time.Time
}

type SamTimeManager struct{}

func (SamTimeManager) Now() time.Time {
	return time.Now()
}
