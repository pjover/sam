package util

import "time"

type Clock interface {
	Now() time.Time
}

type SamClock struct{}

func (SamClock) Now() time.Time { return time.Now() }
