package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/manue1/myhttp/internal/result"
	"github.com/manue1/myhttp/pkg/request"
)

func main() {
	start := time.Now()

	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	log.Println("parallelCount: ", *parallelCount)

	if len(flag.Args()) == 0 {
		log.Println("Please provide URLs as an argument")
		os.Exit(0)
	}

	reqClient := result.NewClient()
	output := request.NewOutput()
	request.StartBatch(flag.Args(), *parallelCount, reqClient, output)

	end := time.Now()
	diff := end.Sub(start)
	log.Printf("time in total %f seconds", diff.Seconds())
}
