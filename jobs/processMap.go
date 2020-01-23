package jobs

//ProcessMap is used to maintain task definitions defined by keys
var ProcessMap map[string]func() = map[string]func(){
	"UPDATE_FEEDS":      UpdateFeedsJob,
	"CHECK_FOR_FEEDS":   AddFeedsJob,
	"CHECK_FOR_PROCESS": CheckForProcesses,
	"DUMP_FEEDS":        DumpFeedsJob,
}
