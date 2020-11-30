package main

import (
	"fmt"
	"poseidon/db"
	"poseidon/runner"
)

func main() {
	fmt.Println("Lets go")
	db.InitializeDatabase()
	runner.InitializeJobMaster()
	select {}
}
