package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nhtuan0700/go-worker/worker"
)

type TestJob struct{}

func NewTestJob() TestJob {
	return TestJob{}
}

func (tj TestJob) Run() error {
	time.Sleep(5 * time.Second)
	fmt.Println("Test job run background!!!")
	return nil
}
func main() {
	jobWorker := worker.NewWorker()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		testJob := NewTestJob()
		jobWorker.Enqueue(testJob)
		w.Write([]byte("Test successfully"))
	})

	addr := ":8081"
	fmt.Println("Server is listening at: ", addr)
	http.ListenAndServe(addr, nil)
}
