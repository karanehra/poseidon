package model

//Job defines a function to be executed at some times
type Job struct {
	Execute func()
	Name    string
}
