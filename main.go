package main

import (
	"flag"
	"log"
	"os"

	"github.com/manue1/myhttp/internal/result"
	"github.com/manue1/myhttp/pkg/job"
)

func main() {
	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Println("Please provide URLs as an argument")
		os.Exit(0)
	}

	reqClient := result.NewClient()
	output := job.NewOutput()
	job.Start(flag.Args(), *parallelCount, reqClient, output)
}
