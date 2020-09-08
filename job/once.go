/*
 * Author: fasion
 * Created time: 2019-05-10 16:28:38
 * Last Modified by: fasion
 * Last Modified time: 2020-09-08 16:42:49
 */

package job

import (
	"fmt"
)

var JOB_ERROR_REENTRY = fmt.Errorf("job reentry")
var JOB_GROUP_FULL = fmt.Errorf("job group full")

type OnceJob interface {
	Job
	Process()
}

type OnceJobRunner struct {
	mutexChan chan struct{}
	enterChan chan struct{}
	exitChan  chan struct{}

	job OnceJob
}

func NewOnceJobRunner(job OnceJob) (*OnceJobRunner, error) {
	return &OnceJobRunner{
		mutexChan: make(chan struct{}, 1),
		enterChan: make(chan struct{}),
		exitChan:  make(chan struct{}),

		job: job,
	}, nil
}

func (runner *OnceJobRunner) RunForever() {
	runner.job.Process()
}

func (runner *OnceJobRunner) Start() error {
	runner.mutexChan <- struct{}{}

	select {
	case <-runner.enterChan:
		<-runner.mutexChan
		return JOB_ERROR_REENTRY
	default:
	}

	ctx := runner.job.GetContext()
	jg := runner.job.GetJobGroup()

	// can not enter job group
	if !jg.Enter(ctx) {
		<-runner.mutexChan
		return JOB_GROUP_FULL
	}

	go func() {
		close(runner.enterChan)
		<-runner.mutexChan

		defer close(runner.exitChan)
		defer jg.Exit()

		// TODO
		// what about panic
		runner.RunForever()
	}()

	return nil
}

func (runner *OnceJobRunner) Shutdown() error {
	runner.job.Cancel()
	return nil
}

func (runner *OnceJobRunner) Join() error {
	select {
	case <-runner.exitChan:
	}

	return nil
}

func (runner *OnceJobRunner) Stop() error {
	if err := runner.Shutdown(); err != nil {
		return err
	}

	if err := runner.Join(); err != nil {
		return err
	}

	return nil
}
