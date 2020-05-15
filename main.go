package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/manue1/myhttp/pkg/handler"
)

func main() {
	start := time.Now()

	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	log.Println("parallelCount: ", *parallelCount)

	var (
		urls     = flag.Args()
		urlCount = len(urls)

		urlChan = make(chan string, urlCount)
		results = make(chan handler.Result, urlCount)
		done    = make(chan struct{})
	)

	log.Println("urls: ", urls)

	go allocateUrls(urls, urlChan)

	go printResults(done, results)

	var workerCount int
	if urlCount < *parallelCount {
		workerCount = urlCount
	} else {
		workerCount = *parallelCount
	}

	createWorkerPool(workerCount, urlChan, results)
	<-done

	end := time.Now()
	diff := end.Sub(start)
	log.Printf("time in total %f seconds", diff.Seconds())
}

func allocateUrls(urls []string, urlChan chan string) {
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)
}

func printResults(done chan struct{}, results chan handler.Result) {
	for r := range results {
		log.Printf(r.String())
	}
	done <- struct{}{}
}

func createWorkerPool(workerCount int, urls chan string, results chan handler.Result) {
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, urls, results)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup, urls chan string, results chan handler.Result) {
	for url := range urls {
		r := handler.GetHashedResponse(url)
		results <- r
	}
	wg.Done()
}
