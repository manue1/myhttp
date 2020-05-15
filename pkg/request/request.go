package request

import (
	"log"
	"sync"

	"github.com/manue1/myhttp/internal/result"
)

func StartBatch(urls []string, parallelCount int) {
	log.Println("urls: ", urls)

	var (
		urlCount    = len(urls)
		workerCount = parallelCount

		urlChan = make(chan string, urlCount)
		results = make(chan result.Page, urlCount)
		done    = make(chan struct{})
	)

	go allocateUrls(urls, urlChan)
	go printResults(done, results)

	if urlCount < parallelCount {
		workerCount = urlCount
	}

	createWorkerPool(workerCount, urlChan, results)

	<-done
}

func allocateUrls(urls []string, urlChan chan string) {
	for _, url := range urls {
		urlChan <- url
	}

	close(urlChan)
}

func printResults(done chan struct{}, results chan result.Page) {
	for r := range results {
		log.Printf(r.String())
	}

	done <- struct{}{}
}

func createWorkerPool(workerCount int, urls chan string, results chan result.Page) {
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, urls, results)
	}

	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup, urls chan string, results chan result.Page) {
	resultClient := result.NewClient()
	for url := range urls {
		r := resultClient.Get(url)
		results <- r
	}

	wg.Done()
}
