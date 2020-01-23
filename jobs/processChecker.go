package jobs

import (
	"fmt"
	"juno/database"
	"poseidon/logger"

	"github.com/karanehra/schemas"
)

//CheckForProcesses finds any processes in the database
func CheckForProcesses() {
	logger := &logger.Logger{}
	logger.INFO("Starting check for process job")
	logger.DepthIn()
	process := schemas.GetNewProcess(database.DB)
	logger.INFO(fmt.Sprintf("Found %v processe", process))
	logger.DepthOut()
	logger.SUCCESS("Job finished")
}
