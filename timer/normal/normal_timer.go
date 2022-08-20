package normal

import (
	"github.com/phuhao00/broker/timer"
	"sync"
	"time"
)

type TimerNormal struct {
	tickTime  time.Duration
	callbacks sync.Map
}

func (t *TimerNormal) AddCallBack(info *timer.CallBackInfo) {
	t.callbacks.Store(info, struct{}{})
}

func (t *TimerNormal) DelCallBack(info *timer.CallBackInfo) {
	t.callbacks.Delete(info)
}

func (t *TimerNormal) Process() {
	t.callbacks.Range(func(key, value any) bool {
		cbInfo := key.(*timer.CallBackInfo)
		cbInfo.Do()
		return true
	})
}

func (t *TimerNormal) Run() {
	go func() {
		tick := time.NewTicker(t.tickTime)
		defer func() {
			tick.Stop()
			err := recover()
			if err != nil {

			}
		}()
		for {
			select {
			case <-tick.C:
				t.Process()
			}
		}
	}()
}
