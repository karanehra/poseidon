package jobs

import "poseidon/model"

//SpawnNewMaster creates a new master of a maximum worker capacity
func SpawnNewMaster(numberOfWorkers int) *model.Master {
	return &model.Master{
		Workers:     []*model.Worker{},
		WorkerLimit: numberOfWorkers,
	}
}
