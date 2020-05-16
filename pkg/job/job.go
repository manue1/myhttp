package job

import (
	"sync"

	"github.com/manue1/myhttp/internal/result"
)

// Start creates the worker pool, waits for the workers to finish
// and passes their results to the provided output
func Start(
	urls []string,
	parallelCount int,
	reqClient result.Client,
	output outputPrinter,
) {
	var (
		urlCount    = len(urls)
		workerCount = parallelCount

		urlChan = make(chan string, urlCount)
		results = make(chan result.Page, urlCount)
		done    = make(chan struct{})
	)

	go allocateUrls(urls, urlChan)
	go output.Print(done, results)

	if urlCount < parallelCount {
		workerCount = urlCount
	}

	createWorkerPool(reqClient, workerCount, urlChan, results)

	<-done
}

// allocateUrls allocates the urls in the channel
func allocateUrls(urls []string, urlChan chan string) {
	for _, url := range urls {
		urlChan <- url
	}

	close(urlChan)
}

// createWorkerPool creates the amount of needed workers and waits for them to finish
func createWorkerPool(reqClient result.Client, workerCount int, urls chan string, results chan result.Page) {
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, reqClient, urls, results)
	}

	wg.Wait()
	close(results)
}

// worker retrieves the result from the client and puts them in the result channel
func worker(wg *sync.WaitGroup, reqClient result.Client, urls chan string, results chan result.Page) {
	for url := range urls {
		r := reqClient.Get(url)

		results <- r
	}

	wg.Done()
}
