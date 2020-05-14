package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	parallelCount := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	log.Println("parallelCount: ", *parallelCount)
	log.Println("urls: ", flag.Args())

	for _, url := range flag.Args() {
		sanitizedUrl := sanitizeProtocol(url)

		resp, err := doRequest(sanitizedUrl)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s %s", sanitizedUrl, computeHash(resp))
	}
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
