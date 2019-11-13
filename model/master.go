package model

//Master manages mulitple workers
type Master struct {
	Workers     []*Worker
	WorkerLimit int
}

//AddWorker spawns a new worker for the master
func (master *Master) AddWorker() {
	worker := &Worker{}
	master.Workers = append(master.Workers, worker)
}

//AddJob sends a new job to the master
func (master *Master) AddJob(job *Job) {
	freeWorker := master.GetFreeWorker()
	freeWorker.Enqueue(job)
}

//GetFreeWorker return a refrence the first least busy worker
func (master *Master) GetFreeWorker() *Worker {
	var freeWorkerIndex int
	for i := range master.Workers {
		if master.Workers[i] == nil {
			freeWorkerIndex = i
			break
		}
	}
	if freeWorkerIndex == 0 {
		master.AddWorker()
		return master.Workers[0]
	}
	return master.Workers[freeWorkerIndex]
}
