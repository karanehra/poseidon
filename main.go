package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func main() {
	const PORT = 3002
	router := mux.NewRouter()
	color.HiGreen("Starting POSEIDON on PORT:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
