/*
 * Author: fasion
 * Created time: 2019-05-10 16:28:38
 * Last Modified by: fasion
 * Last Modified time: 2019-09-10 13:47:17
 */

package job

type RepeatJob interface {
	Job

	Prepare() (bool)
	Process() (bool)
	Cleanup()
}

type RepeatJobRunner struct {
	JobRunner
	Job

	job 			RepeatJob
}

func NewRepeatJobRunner(job RepeatJob) (*RepeatJobRunner, error) {
	runner := RepeatJobRunner{
		Job: job,
		job: job,
	}

	onceRunner, err := NewOnceJobRunner(&runner)
	if err != nil {
		return nil, err
	}

	runner.JobRunner = onceRunner

	return &runner, nil
}

func (self *RepeatJobRunner) Process() {
	// call prepare
	if self.job.IsCanceled() {
		return
	}
	if !self.job.Prepare() {
		return
	}

	// call process repeatly
	for {
		if self.job.IsCanceled() {
			break
		}

		if !self.job.Process() {
			break
		}
	}

	// call clean up
	self.job.Cleanup()
}
