package timer

import "time"

type Category int

const (
	Daily Category = iota + 1
	Weekly
	Monthly
	Once
)

type CallBackInfo struct {
	Category     Category
	interval     time.Duration
	lastCallTime time.Time
	fn           func()
}

func (cb *CallBackInfo) Do() {
	cb.fn()
	cb.lastCallTime = time.Now()
}
