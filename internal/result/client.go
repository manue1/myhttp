package result

import "net/http"

type (
	Client struct {
		http httpClient
	}

	httpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewClient() Client {
	return Client{
		http: &http.Client{},
	}
}
