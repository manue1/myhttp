package main

import (
	"flag"
	"log"
)

func main() {
	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	log.Println("parallelCount: ", *parallelCount)
	log.Println("urls: ", flag.Args())

	// @ToDo in parallel:
	// - make http request
	// - hash the response body
	// - print address hash (prepend http:// if not stated as part of arg)
}
