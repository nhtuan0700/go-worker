package worker

import (
	"fmt"
)

type Job interface {
	Run() error
}

type Worker struct {
	Jobs chan Job
}

func (w *Worker) Enqueue(job Job) {
	select {
	case w.Jobs <- job:
	default:
		// Log or handle full channel case
		fmt.Println("Job queue is full")
	}
}

func NewWorker() *Worker {
	jobs := make(chan Job, 10)
	worker := &Worker{
		Jobs: jobs,
	}

	go func() {
		for job := range worker.Jobs {
			err := job.Run()
			if err != nil {
				fmt.Println("Job error:", err)
			}
		}
	}()

	return worker
}
