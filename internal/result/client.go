package result

import "net/http"

type (
	Client struct {
		Http HttpClient
	}

	HttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewClient() Client {
	return Client{
		Http: &http.Client{},
	}
}
