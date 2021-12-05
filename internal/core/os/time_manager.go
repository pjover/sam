package os

import "time"

type TimeManager interface {
	Now() time.Time
}

type timeManager struct {
}

func NewTimeManager() TimeManager {
	return timeManager{}
}

func (timeManager) Now() time.Time {
	return time.Now()
}
