package model

//Worker runs a job queue
type Worker struct {
	IsBusy   bool
	JobQueue []func()
}
