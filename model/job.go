package model

//Job defines a function to be executed at some times
type Job struct {
	Executer      func()
	Name          string
	ParentChannel chan int
}

//Run starts the job in a goroutine
func (job *Job) Run() {
	go job.Executer()
	job.ParentChannel <- 1
}
