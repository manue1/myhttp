package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Client struct{}

func (m *Client) Do(req *http.Request) (*http.Response, error) {
	content := getMockContent(req.URL.String())

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(content))),
	}, nil
}
