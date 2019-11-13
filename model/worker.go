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
	worker.JobQueue = append(worker.JobQueue, job)
	fmt.Printf("Enqued: len = %v \n", len(worker.JobQueue))
	if !worker.IsBusy {
		worker.Start()
	}
}

//Start starts worker execution
func (worker *Worker) Start() {
	go func() {
		for {
			select {
			case <-worker.Channel:
				fmt.Println("Got it")
				job := worker.getJobToExecute()
				if job != nil {
					job.Run(worker.Channel)
				}
			}
		}
	}()
	worker.IsBusy = true
	worker.Channel = make(chan int)
	worker.getJobToExecute().Run(worker.Channel)
}

func (worker *Worker) getJobToExecute() *Job {
	if len(worker.JobQueue) > 0 {
		currentTask := worker.JobQueue[0]
		if len(worker.JobQueue) <= 1 {
			worker.JobQueue = []*Job{}
		} else {
			worker.JobQueue = worker.JobQueue[1:]
		}
		return currentTask
	}
	worker.IsBusy = false
	return nil
}
