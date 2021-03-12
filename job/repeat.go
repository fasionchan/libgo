/*
 * Author: fasion
 * Created time: 2019-05-10 16:28:38
 * Last Modified by: fasion
 * Last Modified time: 2020-09-23 20:02:29
 */

package job

type RepeatJob interface {
	Job

	Prepare() bool
	Process() bool
	CleanUp()
}

type RepeatJobRunner struct {
	JobRunner
	RepeatJob

	// job RepeatJob
}

func NewRepeatJobRunner(job RepeatJob) (*RepeatJobRunner, error) {
	runner := RepeatJobRunner{
		RepeatJob: job,
		// job: job,
	}

	onceRunner, err := NewOnceJobRunner(&runner)
	if err != nil {
		return nil, err
	}

	runner.JobRunner = onceRunner

	return &runner, nil
}

func (runner *RepeatJobRunner) Process() {
	// call prepare
	if runner.RepeatJob.IsCanceled() {
		return
	}
	if !runner.RepeatJob.Prepare() {
		return
	}

	// call process repeatly
	for {
		if runner.RepeatJob.IsCanceled() {
			break
		}

		if !runner.RepeatJob.Process() {
			break
		}
	}

	// call clean up
	runner.RepeatJob.CleanUp()
}
