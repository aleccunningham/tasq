package job

import "time"

type JobState int

const (
	Ok JobState = iota
	Failed
	Pending
)

// Job wraps function calls
type Job struct {
	id int
	fn func()

	startTime time.Duration
	endTime   time.Duration
	status    JobState
}

func (j *Job) GetID() int {
	return j.id
}

type JobResult struct {
	name  string
	value []byte
}
