/*
 * Author: fasion
 * Created time: 2019-05-10 16:40:29
 * Last Modified by: fasion
 * Last Modified time: 2020-04-23 09:10:33
 */

package job

import (
	"context"
	"sync"
	"time"
)

type JobRunner interface {
	RunForever()
	Start() error
	Shutdown() error
	Join() error
	Stop() error
}

type JobGroup struct {
	tokens    chan struct{}
	wg        sync.WaitGroup
	enterWait time.Duration
}

func NewJobGroup(tokens int, enterWait time.Duration) *JobGroup {
	return &JobGroup{
		tokens:    make(chan struct{}, tokens),
		enterWait: enterWait,
	}
}

func (self *JobGroup) Enter(ctx context.Context) bool {
	if self == nil {
		return true
	}

	if self.tokens == nil {
		self.wg.Add(1)
		return true
	}

	if self.enterWait == 0 {
		select {
		case self.tokens <- struct{}{}:
			self.wg.Add(1)
			return true
		default:
			return false
		}
	}

	select {
	case self.tokens <- struct{}{}:
		self.wg.Add(1)
		return true
	case <-ctx.Done():
		return false
	case <-time.After(self.enterWait):
		return false
	}
}

func (self *JobGroup) Exit() {
	if self == nil {
		return
	}

	self.wg.Done()

	if self.tokens != nil {
		<-self.tokens
	}
}

func (self *JobGroup) Wait() {
	if self == nil {
		return
	}

	self.wg.Wait()
}

type Job interface {
	GetJobGroup() *JobGroup
	GetContext() context.Context
	Cancel()
	IsCanceled() bool
	OnError(error)
	GetLastError() error
}

type BaseJob struct {
	ident  string
	ctx    context.Context
	cancel context.CancelFunc

	lastErr error
}

func NewBaseJob(ident string) BaseJob {
	ctx, cancel := context.WithCancel(context.Background())
	return BaseJob{
		ident:  ident,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (job *BaseJob) GetJobIdent() string {
	return job.ident
}

func (self *BaseJob) GetJobGroup() *JobGroup {
	return nil
}

func (self *BaseJob) GetContext() context.Context {
	return self.ctx
}

func (self *BaseJob) Cancel() {
	self.cancel()
}

func (self *BaseJob) IsCanceled() bool {
	select {
	case <-self.ctx.Done():
		return true
	default:
		return false
	}
}

func (base *BaseJob) Sleep(interval time.Duration) bool {
	select {
	case <-base.GetContext().Done():
		return false
	case <-time.After(interval):
		return true
	}
}

func (self *BaseJob) OnError(err error) {
	self.lastErr = err
}

func (self *BaseJob) GetLastError() error {
	return self.lastErr
}
