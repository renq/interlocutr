package clock

import "time"

type Clock struct {
	currentTime *time.Time
}

func (c Clock) Now() time.Time {
	if c.currentTime != nil {
		return *c.currentTime
	}

	return time.Now().UTC()
}

func NewClock() *Clock {
	return &Clock{}
}

func (c *Clock) FreezeTime(time time.Time) {
	c.currentTime = &time
}
