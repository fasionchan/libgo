/*
 * Author: fasion
 * Created time: 2020-11-18 15:17:36
 * Last Modified by: fasion
 * Last Modified time: 2020-11-18 16:38:51
 */

package job

import (
	"context"
	"time"
)

type Restrictor struct {
	BaseJob
	JobRunner
	batchSize int
	tickets   chan interface{}
}

func NewRestrictor(ident string, capacity, batchSize int, batchInterval time.Duration) (*Restrictor, error) {
	r := Restrictor{
		BaseJob:   NewBaseJob(ident),
		batchSize: batchSize,
		tickets:   make(chan interface{}, capacity),
	}

	runner, err := NewPeriodJobRunner(&r, batchInterval, 0, false, false)
	if err != nil {
		return nil, err
	}

	r.JobRunner = runner

	return &r, nil
}

func (r Restrictor) Tickets() <-chan interface{} {
	return r.tickets
}

func (r Restrictor) WaitForTicket(ctx context.Context) bool {
	if ctx == nil {
		<-r.tickets
		return true
	}

	select {
	case <-ctx.Done():
		return false
	case <-r.tickets:
		return true
	}
}

func (r Restrictor) Prepare() bool {
	return true
}

func (r *Restrictor) Process() bool {
	for i := 0; i < r.batchSize; i++ {
		r.tickets <- nil
	}

	return true
}

func (r Restrictor) CleanUp() {
}
