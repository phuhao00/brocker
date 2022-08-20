package normal

import (
	"fmt"
	"github.com/phuhao00/broker/timerassistant"
	"sync"
	"time"
)

type TimerNormalAssistant struct {
	tickTime  time.Duration
	callbacks sync.Map
}

func NewTimerNormalAssistant(tickTime time.Duration) *TimerNormalAssistant {
	return &TimerNormalAssistant{
		tickTime: tickTime,
	}
}

func (t *TimerNormalAssistant) AssertOwner(owner timerassistant.Owner) {

}

func (t *TimerNormalAssistant) AddCallBack(info *timerassistant.CallInfo) {
	t.callbacks.Store(info, struct{}{})
}

func (t *TimerNormalAssistant) DelCallBack(info *timerassistant.CallInfo) {
	t.callbacks.Delete(info)
}

func (t *TimerNormalAssistant) Process() {
	t.callbacks.Range(func(key, value any) bool {
		cbInfo := key.(*timerassistant.CallInfo)
		cbInfo.NotifyOwner()
		return true
	})
}

func (t *TimerNormalAssistant) Loop() {
	go func() {
		tick := time.NewTicker(t.tickTime)
		defer func() {
			tick.Stop()
			err := recover()
			if err != nil {
				fmt.Println(err)
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
