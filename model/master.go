package model

//Master manages mulitple workers
type Master struct {
	Workers []*Worker
}

//AddWorker spawns a new worker for the master
func (master *Master) AddWorker() {
	worker := &Worker{}
	master.Workers = append(master.Workers, worker)
}

//AddJob sends a new job to the master
func (master *Master) AddJob(job *Job) {
	//To write
}
