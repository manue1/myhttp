package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type result struct {
	url          string
	hashResponse string
}

func main() {
	start := time.Now()

	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	log.Println("parallelCount: ", *parallelCount)

	var (
		urls     = flag.Args()
		urlCount = len(urls)

		urlChan = make(chan string, urlCount)
		results = make(chan result, urlCount)
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

func printResults(done chan struct{}, results chan result) {
	for r := range results {
		log.Printf("%s %s", r.url, r.hashResponse)
	}
	done <- struct{}{}
}

func createWorkerPool(workerCount int, urls chan string, results chan result) {
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, urls, results)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup, urls chan string, results chan result) {
	for url := range urls {
		r := getHashedResponse(url)
		results <- r
	}
	wg.Done()
}

func getHashedResponse(argUrl string) result {
	url := sanitizeProtocol(argUrl)

	resp, err := doRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	hashResponse := computeHash(resp)

	return result{url, hashResponse}
}

func doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func sanitizeProtocol(url string) string {
	var (
		protocol  = "http://"
		sanitized = url
	)

	if !strings.HasPrefix(url, protocol) {
		sanitized = protocol + url
	}

	return sanitized
}

func computeHash(response []byte) string {
	h := md5.New()
	h.Write(response)

	return hex.EncodeToString(h.Sum(nil))
}
