package throttle

import (
	"fmt"
	"time"
)

type Throttle struct {
	tAfter uint64        // trottling should start after tAfter events/messages
	tDelay time.Duration //once trottling started, wait for tDelay for each event/message send
	ch     chan int      // use this channel to notify
}

func Init(tAfter uint64, tDelay time.Duration) *Throttle {
	ch := make(chan int, tAfter)
	t := &Throttle{tAfter: tAfter, tDelay: tDelay, ch: ch}
	return t

}
func InitAndStartThrottle(tAfter uint64, tDelay time.Duration) *Throttle {
	t := Init(tAfter, tDelay)
	t.Start()
	return t
}

func (t *Throttle) Start() {
	go func() {
		var i uint64
		for i = 0; i < t.tAfter; i += 1 {
			t.ch <- 1
		}
		for {
			fmt.Println("push")
			t.ch <- 1

			<-time.After(t.tDelay)

		}
	}()

}

func (t *Throttle) Check() int {
	return <-t.ch
}
