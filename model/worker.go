package model

//Worker runs a job queue
type Worker struct {
	IsBusy   bool
	JobQueue []*Job
}

//Enqueue adds a job to the queue
func (worker *Worker) Enqueue(job *Job) {
	worker.JobQueue = append(worker.JobQueue, job)
}
