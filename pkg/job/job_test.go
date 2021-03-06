package job

import (
	"testing"

	"github.com/manue1/myhttp/internal/result"
	"github.com/manue1/myhttp/test/mocks"
)

func TestStartJob(t *testing.T) {
	testCases := []struct {
		scenario      string
		urls          []string
		parallelCount int
	}{
		{
			scenario:      "Test success - custom parallel value",
			urls:          []string{"http://adjust.com", "http://facebook.com"},
			parallelCount: 1,
		},
		{
			scenario:      "Test success - default parallel value",
			urls:          []string{"adjust.com", "facebook.com", "google.com", "twitter.com"},
			parallelCount: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			outputMock := outputMock{t}
			clientMock := result.Client{
				Http: &mocks.Client{},
			}

			Start(tc.urls, tc.parallelCount, clientMock, outputMock)
		})
	}
}

type outputMock struct {
	t *testing.T
}

func (o outputMock) Print(done chan struct{}, results chan result.Page) {
	for r := range results {
		expectedHash := mocks.GetMockMD5(r.URL)
		if expectedHash != r.HashResponse {
			o.t.Errorf("unexpected md5 hash: got %s want %s",
				r.HashResponse, expectedHash)
		}
	}

	done <- struct{}{}
}
