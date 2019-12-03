package main

import (
	"poseidon/jobs"
)

func main() {
	jobs.LaunchRunner()
	select {}
}
