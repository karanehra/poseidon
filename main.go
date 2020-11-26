package main

import (
	"fmt"
	"poseidon/db"
)

func main() {
	fmt.Println("Lets go")
	db.InitializeDatabase()
	//This select keeps the process running
	checkJobs()
	select {}
}
