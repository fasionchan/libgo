/*
 * Author: fasion
 * Created time: 2019-05-13 13:50:47
 * Last Modified by: fasion
 * Last Modified time: 2019-10-31 10:50:34
 */

package job

import (
	"time"
)

type PeriodJob interface {
	Job

	Prepare() bool
	Process() bool
	Cleanup()
}

type PeriodJobRunner struct {
	JobRunner
	Job

	interval    time.Duration
	offset      time.Duration
	align       bool
	strictAlign bool

	job PeriodJob

	nextTickTime time.Time
}

func NewPeriodJobRunner(job PeriodJob, interval time.Duration, offset time.Duration, align, strictAlign bool) (*PeriodJobRunner, error) {
	if !align {
		offset = 0
	}

	runner := PeriodJobRunner{
		Job:         job,
		interval:    interval,
		offset:      offset,
		align:       align,
		strictAlign: strictAlign,
		job:         job,
	}

	onceRunner, err := NewOnceJobRunner(&runner)
	if err != nil {
		return nil, err
	}

	runner.JobRunner = onceRunner

	return &runner, nil
}

func (self *PeriodJobRunner) Process() {
	// call prepare
	if self.job.IsCanceled() {
		return
	}
	if !self.job.Prepare() {
		return
	}

	curTime := time.Now()

	// initialize next executing time
	if self.align && self.strictAlign {
		self.nextTickTime = alignNextTime(curTime, self.interval, self.offset)
	} else {
		self.nextTickTime = curTime
	}

PERIOD_RUNNER_LOOP:
	for {
		curTime = time.Now()

		if curTime.Before(self.nextTickTime) {
			// not right now, waiting
			waitDuration := self.nextTickTime.Sub(curTime)
			if self.align && waitDuration > self.interval {
				self.nextTickTime = alignNextTime(curTime, self.interval, self.offset)
				continue
			}

			select {
			case <-time.After(waitDuration):
			case <-self.job.GetContext().Done():
				break PERIOD_RUNNER_LOOP
			}
		} else {
			if self.job.IsCanceled() {
				break
			}
			if !self.job.Process() {
				break
			}

			// calculate next executing time
			if self.align {
				self.nextTickTime = alignNextTime(curTime, self.interval, self.offset)
			} else {
				self.nextTickTime = self.nextTickTime.Add(self.interval)
			}
		}
	}

	// call clean up
	self.job.Cleanup()
}

func alignNextTimeSmart(base time.Time, interval time.Duration, offset time.Duration) time.Time {
	next := base.Truncate(interval).Add(interval).Add(offset)
	return base.Add(next.Sub(base))
}

func alignNextTime(base time.Time, interval time.Duration, offset time.Duration) time.Time {
	return base.Truncate(interval).Add(interval).Add(offset)
}
