package request

import (
	"testing"

	"github.com/manue1/myhttp/internal/result"
	"github.com/manue1/myhttp/test/mocks"
)

func TestStartBatch(t *testing.T) {
	testCases := []struct {
		scenario      string
		urls          []string
		parallelCount int
	}{
		{
			scenario:      "Test success",
			urls:          []string{"http://adjust.com", "http://facebook.com"},
			parallelCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			outputMock := outputMock{t}
			clientMock := result.Client{
				Http: &mocks.Client{},
			}

			StartBatch(tc.urls, tc.parallelCount, clientMock, outputMock)
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
