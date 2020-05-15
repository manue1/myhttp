package result

import "net/http"

type (
	Client struct {
		http HttpClient
	}

	HttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewClient() Client {
	return Client{
		http: &http.Client{},
	}
}
