package result

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	URL          string
	HashResponse string
}

func (r Page) String() string {
	return fmt.Sprintf("%s %s", r.URL, r.HashResponse)
}

func (c Client) Get(argUrl string) Page {
	url := sanitizeProtocol(argUrl)

	resp, err := c.doRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	hashResponse := computeHash(resp)

	return Page{URL: url, HashResponse: hashResponse}
}

func (c Client) doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Http.Do(req)
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
