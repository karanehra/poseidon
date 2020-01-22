package jobs

import (
	"fmt"
	"poseidon/db"
	"poseidon/logger"

	"github.com/karanehra/schemas"
)

//CheckForProcesses finds any processes in the database
func CheckForProcesses() {
	logger := &logger.Logger{}
	logger.INFO("Starting check for process job")
	logger.DepthIn()
	processes, err := schemas.GetAllProcesses(db.DB)
	if err != nil {
		logger.ERROR("Job failed due to error")
		return
	}
	logger.INFO(fmt.Sprintf("Found %v processes", len(processes)))
	logger.DepthOut()
	logger.SUCCESS("Job finished")
}
