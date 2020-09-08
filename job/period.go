/*
 * Author: fasion
 * Created time: 2019-05-13 13:50:47
 * Last Modified by: fasion
 * Last Modified time: 2020-09-08 16:46:49
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

func (runner *PeriodJobRunner) Process() {
	// call prepare
	if runner.job.IsCanceled() {
		return
	}
	if !runner.job.Prepare() {
		return
	}

	curTime := time.Now()

	// initialize next executing time
	if runner.align && runner.strictAlign {
		runner.nextTickTime = alignNextTime(curTime, runner.interval, runner.offset)
	} else {
		runner.nextTickTime = curTime
	}

PERIOD_RUNNER_LOOP:
	for {
		curTime = time.Now()

		if curTime.Before(runner.nextTickTime) {
			// not right now, waiting
			waitDuration := runner.nextTickTime.Sub(curTime)
			if runner.align && waitDuration > runner.interval {
				runner.nextTickTime = alignNextTime(curTime, runner.interval, runner.offset)
				continue
			}

			select {
			case <-time.After(waitDuration):
			case <-runner.job.GetContext().Done():
				break PERIOD_RUNNER_LOOP
			}
		} else {
			if runner.job.IsCanceled() {
				break
			}
			if !runner.job.Process() {
				break
			}

			// calculate next executing time
			if runner.align {
				runner.nextTickTime = alignNextTime(curTime, runner.interval, runner.offset)
			} else {
				runner.nextTickTime = runner.nextTickTime.Add(runner.interval)
			}
		}
	}

	// call clean up
	runner.job.Cleanup()
}

func alignNextTimeSmart(base time.Time, interval time.Duration, offset time.Duration) time.Time {
	next := base.Truncate(interval).Add(interval).Add(offset)
	return base.Add(next.Sub(base))
}

func alignNextTime(base time.Time, interval time.Duration, offset time.Duration) time.Time {
	return base.Truncate(interval).Add(interval).Add(offset)
}
