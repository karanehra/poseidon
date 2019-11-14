package model

//Job defines a function to be executed at some times
type Job struct {
	Executer      func(interface{})
	Name          string
	ParentChannel *chan int
	Payload       interface{}
}

//Run starts the job in a goroutine
func (job *Job) Run(channel chan int) {
	job.Executer(job.Payload)
}

//AddPayload is used to define the data to inject into the job executor
func (job *Job) AddPayload(payload interface{}) {
	job.Payload = payload
}

//AddPayloadAndReturn is used to define the data to inject into the job executor.
// Returns a reference to the Job
func (job *Job) AddPayloadAndReturn(payload interface{}) *Job {
	job.Payload = payload
	return job
}
