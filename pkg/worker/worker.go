package worker

import (
	"time"
)

type Worker struct {
	name    string
	jq      JobQueue
	rq      ResultQueue
	lastJob *Job
	debug   bool
	done    chan bool
}

func (w *Worker) Name() (string, error) {
	if w.name == "" {
		return
	}

	return w.name
}

func (w *Worker) Results() *ResultQueue {
	return w.rq
}

func (w *Worker) Done() {
	w.done <- true
}

func (w *Worker) Run() {
	w.lastJob = w.jq.GetTask()
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case ticker.C:
			w.Run()
		}
	}
}
