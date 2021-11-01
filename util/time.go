package util

import "time"

type Timer interface {
	Now() time.Time
}

type SamTimer struct{}

func (SamTimer) Now() time.Time {
	return time.Now()
}
