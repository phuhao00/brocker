package timerassistant

import "time"

type CallCategory interface {
	ShouldCall() bool
	SetLastCallTime(int642 int64)
}

type Once struct {
	Hour         int
	Min          int
	Sec          int
	lastCallTime int64
}

func (o *Once) ShouldCall() bool {
	if o.lastCallTime != 0 {
		return false
	}
	now := time.Now()
	hour, min, sec := now.Local().Clock()
	if (hour*3600 + min*60 + sec) >= (o.Hour*3600 + o.Min*60 + o.Sec) {
		return true
	}
	return false
}
func (o *Once) SetLastCallTime(lastCallTime int64) {
	o.lastCallTime = lastCallTime
}

type Daily struct {
	Hour         int64
	Min          int64
	Sec          int64
	lastCallTime int64
}

func (d *Daily) ShouldCall() bool {
	now := time.Now().Local()
	hour, min, sec := now.Clock()
	nowSec := int64(hour*3660) + int64(min*60) + int64(sec)

	if d.lastCallTime == 0 {
		if nowSec >= (d.Hour*3600 + d.Hour*60 + d.Sec) {
			return true
		}
	} else {
		if nowSec >= (d.Hour*3600+d.Hour*60+d.Sec) && nowSec-d.lastCallTime >= int64(86400) {
			return true
		}
	}

	return false
}

func (d *Daily) SetLastCallTime(lastCallTime int64) {
	d.lastCallTime = lastCallTime

}

type Weekly struct {
	WeekDay int
	Hour    int
	Min     int
	S       int
}
