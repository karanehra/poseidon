package jobs

import (
	"fmt"
	"juno/database"
	"poseidon/db"
	"poseidon/logger"

	"github.com/karanehra/schemas"
)

//CheckForProcesses finds any processes in the database
func CheckForProcesses() {
	logger := &logger.Logger{}
	logger.INFO("Starting check for process job")
	logger.DepthIn()
	process := schemas.GetNewProcess(database.DB)
	fmt.Println(process)
	if process.Type == "" {
		logger.INFO("No process found")
		return
	}
	logger.INFO(fmt.Sprintf("Found %v process: ", process))
	schemas.UpdateProcessStatus(db.DB, "EXECUTING", process.ID)
	ProcessMap[process.Type]()
	schemas.UpdateProcessStatus(db.DB, "FINISHED", process.ID)
	logger.DepthOut()
	logger.SUCCESS("Job finished")
}
