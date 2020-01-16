/*
 * Author: fasion
 * Created time: 2019-06-25 10:10:13
 * Last Modified by: fasion
 * Last Modified time: 2019-06-26 18:56:34
 */

package job

import (
	"context"
	"time"
)

const (
	SleepDone = iota
	SleepCanceled
	SleepInterrupted
)

type InterruptibleSleep struct {
	ctx context.Context
	duration time.Duration
	sig chan interface{}
	lock chan interface{}
}

func NewInterruptibleSleep(ctx context.Context, d time.Duration) (*InterruptibleSleep) {
	return &InterruptibleSleep{
		ctx: ctx,
		duration: d,
		sig: make(chan interface{}),
		lock: make(chan interface{}, 1),
	}
}

func (self *InterruptibleSleep) Interrupt() {
	self.lock <- 1
	defer func(){ <- self.lock }()

	if self == nil {
		return
	}

	select {
	case <- self.sig:
	default:
		close(self.sig)
	}
}

func (self *InterruptibleSleep) Sleep() (int) {
	if self == nil {
		return SleepDone
	}

	select {
	case <- self.ctx.Done():
		return SleepCanceled
	case <- time.After(self.duration):
		return SleepDone
	case <- self.sig:
		return SleepInterrupted
	}
}
