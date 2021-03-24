/*
 * Author: fasion
 * Created time: 2019-05-10 16:40:29
 * Last Modified by: fasion
 * Last Modified time: 2021-03-24 13:42:12
 */

package job

import (
	"context"
	"sync"
	"time"

	"github.com/fasionchan/libgo/logging"
	"go.uber.org/zap"
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

func (group *JobGroup) Enter(ctx context.Context) bool {
	if group == nil {
		return true
	}

	if group.tokens == nil {
		group.wg.Add(1)
		return true
	}

	if group.enterWait == 0 {
		select {
		case group.tokens <- struct{}{}:
			group.wg.Add(1)
			return true
		default:
			return false
		}
	}

	select {
	case group.tokens <- struct{}{}:
		group.wg.Add(1)
		return true
	case <-ctx.Done():
		return false
	case <-time.After(group.enterWait):
		return false
	}
}

func (group *JobGroup) Exit() {
	if group == nil {
		return
	}

	group.wg.Done()

	if group.tokens != nil {
		<-group.tokens
	}
}

func (group *JobGroup) Wait() {
	if group == nil {
		return
	}

	group.wg.Wait()
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
	*zap.Logger
	ident  string
	ctx    context.Context
	cancel context.CancelFunc

	lastErr error
}

func NewBaseJob(ident string) BaseJob {
	ctx, cancel := context.WithCancel(context.Background())
	return BaseJob{
		Logger: logging.GetLogger().Named(ident),
		ident:  ident,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (job *BaseJob) GetJobIdent() string {
	return job.ident
}

func (job *BaseJob) GetJobGroup() *JobGroup {
	return nil
}

func (job *BaseJob) GetContext() context.Context {
	return job.ctx
}

func (job *BaseJob) Cancel() {
	job.cancel()
}

func (job *BaseJob) IsCanceled() bool {
	select {
	case <-job.ctx.Done():
		return true
	default:
		return false
	}
}

func (job *BaseJob) Sleep(interval time.Duration) bool {
	select {
	case <-job.GetContext().Done():
		return false
	case <-time.After(interval):
		return true
	}
}

func (job *BaseJob) OnError(err error) {
	job.lastErr = err
}

func (job *BaseJob) GetLastError() error {
	return job.lastErr
}
