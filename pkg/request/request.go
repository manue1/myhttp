package request

import (
	"log"
	"sync"

	"github.com/manue1/myhttp/internal/result"
)

var reqClient result.Client

func init() {
	reqClient = result.NewClient()
}

func StartBatch(
	urls []string,
	parallelCount int,
	output printer,
) {
	log.Println("urls: ", urls)

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

	createWorkerPool(workerCount, urlChan, results)

	<-done
}

func allocateUrls(urls []string, urlChan chan string) {
	for _, url := range urls {
		urlChan <- url
	}

	close(urlChan)
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
	for url := range urls {
		r := reqClient.Get(url)
		results <- r
	}

	wg.Done()
}
