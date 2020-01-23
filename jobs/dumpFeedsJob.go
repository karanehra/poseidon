package jobs

import "poseidon/logger"

//DumpFeedsJob gets feeds out of DB and dumps them in a csv
func DumpFeedsJob() {
	logger := logger.Logger{}
	logger.INFO("Starting dump feeds job")
	logger.SUCCESS("Job Done")
}
