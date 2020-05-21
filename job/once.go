/*
 * Author: fasion
 * Created time: 2019-05-10 16:28:38
 * Last Modified by: fasion
 * Last Modified time: 2020-04-23 09:11:42
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

func (self *OnceJobRunner) RunForever() {
	self.job.Process()
}

func (self *OnceJobRunner) Start() error {
	self.mutexChan <- struct{}{}

	select {
	case <-self.enterChan:
		<-self.mutexChan
		return JOB_ERROR_REENTRY
	default:
	}

	ctx := self.job.GetContext()
	jg := self.job.GetJobGroup()

	// can not enter job group
	if !jg.Enter(ctx) {
		<-self.mutexChan
		return JOB_GROUP_FULL
	}

	go func() {
		close(self.enterChan)
		<-self.mutexChan

		defer close(self.exitChan)
		defer jg.Exit()

		// TODO
		// what about panic
		self.RunForever()
	}()

	return nil
}

func (self *OnceJobRunner) Shutdown() error {
	self.job.Cancel()
	return nil
}

func (self *OnceJobRunner) Join() error {
	select {
	case <-self.exitChan:
	}

	return nil
}

func (self *OnceJobRunner) Stop() error {
	if err := self.Shutdown(); err != nil {
		return err
	}

	if err := self.Join(); err != nil {
		return err
	}

	return nil
}
