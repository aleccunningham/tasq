package controller

import "context"

type JobQueue struct {
	ctx        context.Context
	numWorkers int
	rq         ResultQueue
}

func (jq *JobQueue) NumWorkers() int {
	return jq.numWorkers
}

func (jq *JobQueue) AddTask(task, ...args) {}

func (jq *JobQueue) Start() {
	for i := range jq.numWorkers {
		w := worker.New()
		w.daemon = true
		w.start()
	}
}
