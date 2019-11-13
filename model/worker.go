package model

import "fmt"

//Worker runs a job queue
type Worker struct {
	IsBusy   bool
	JobQueue []*Job
	Channel  chan int
}

//Enqueue adds a job to the queue
func (worker *Worker) Enqueue(job *Job) {
	job.ParentChannel = worker.Channel
	worker.JobQueue = append(worker.JobQueue, job)
	if !worker.IsBusy {
		worker.Start()
	}
}

//Start starts worker execution
func (worker *Worker) Start() {
	worker.IsBusy = true
	worker.Channel = make(chan int)
	worker.JobQueue[0].Run()
	go func() {
		for {
			select {
			case <-worker.Channel:
				fmt.Println("Got it")
			}
		}
	}()
}
