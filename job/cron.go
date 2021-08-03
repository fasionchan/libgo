/*
 * Author: fasion
 * Created time: 2021-03-10 09:19:12
 * Last Modified by: fasion
 * Last Modified time: 2021-08-03 11:17:40
 */

package job

import (
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronJob interface {
	RepeatJob
}

type CronJobRunner struct {
	JobRunner
	Job

	schedule cron.Schedule
	job      CronJob

	processAtOnce bool
	lastTickTime  time.Time
	nextTickTime  time.Time
}

func NewCronJobRunner(job CronJob, spec string) (*CronJobRunner, error) {
	schedule, err := cron.ParseStandard(spec)
	if err != nil {
		return nil, err
	}

	runner := CronJobRunner{
		Job:      job,
		schedule: schedule,
		job:      job,
	}

	onceRunner, err := NewOnceJobRunner(&runner)
	if err != nil {
		return nil, err
	}

	runner.JobRunner = onceRunner

	return &runner, nil
}

func (runner *CronJobRunner) WithProcessAtOnce(value bool) *CronJobRunner {
	runner.processAtOnce = value
	return runner
}

func (runner *CronJobRunner) Process() {
	// call prepare
	if runner.job.IsCanceled() {
		return
	}
	if !runner.job.Prepare() {
		return
	}

	// first process at once
	if runner.processAtOnce {
		if runner.job.IsCanceled() {
			return
		}
		if !runner.job.Process() {
			return
		}
	}

	runner.lastTickTime = time.Now()

CRON_RUNNER_LOOP:
	for {
		curTime := time.Now()
		runner.nextTickTime = runner.schedule.Next(runner.lastTickTime)

		runner.Info("CronRunnerScheduling",
			zap.String("Ident", runner.job.GetJobIdent()),
			zap.Time("NextTime", runner.nextTickTime),
		)

		if curTime.Before(runner.nextTickTime) {
			// not right now, waiting
			waitDuration := runner.nextTickTime.Sub(curTime)

			select {
			// case <-runner.rescheduleSignal:
			// 	goto START_SCHEDULING
			case <-time.After(waitDuration):
			case <-runner.job.GetContext().Done():
				break CRON_RUNNER_LOOP
			}

		} else {
			if runner.job.IsCanceled() {
				break
			}
			if !runner.job.Process() {
				break
			}

			runner.lastTickTime = time.Now()
		}
	}

	// call clean up
	runner.job.CleanUp()
}
