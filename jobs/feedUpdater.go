package jobs

import "poseidon/runner"

//UpdateFeeds triggers a feed recheck
func UpdateFeeds() {
	//getting all feeds logic
	feeds := []string{"https://timesofindia.indiatimes.com/rssfeedstopstories.cms"}

	for i := range feeds {
		JobMaster.AddJob()
	}
}
