package result

import (
	"testing"

	"github.com/manue1/myhttp/test/mocks"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		scenario    string
		url         string
		expectedURL string
		fn          func(*testing.T, string, Page)
	}{
		{
			scenario:    "Test success - with already sanitized URL",
			url:         "http://adjust.com",
			expectedURL: "http://adjust.com",
			fn:          testSuccessAlreadySanitizedURL,
		},
		{
			scenario:    "Test success - with unsanitized URL",
			url:         "facebook.com",
			expectedURL: "http://facebook.com",
			fn:          testSuccessUnsanitizedURL,
		},
		{
			scenario:    "Test failure - empty response body",
			url:         "localhost.com",
			expectedURL: "http://localhost.com",
			fn:          testFailureEmptyBody,
		},
	}

	client := Client{
		http: &mocks.Client{},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {

			page := client.Get(tc.url)
			tc.fn(t, tc.expectedURL, page)
		})
	}
}

func testSuccessAlreadySanitizedURL(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.url {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.url, expectedURL)
	}

	expectedHash := mocks.GetMockMD5(expectedURL)
	if expectedHash != actualPage.hashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.hashResponse, expectedHash)
	}
}

func testSuccessUnsanitizedURL(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.url {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.url, expectedURL)
	}

	expectedHash := mocks.GetMockMD5(expectedURL)
	if expectedHash != actualPage.hashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.hashResponse, expectedHash)
	}
}

func testFailureEmptyBody(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.url {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.url, expectedURL)
	}

	emptyResponseHash := "d41d8cd98f00b204e9800998ecf8427e"
	if emptyResponseHash != actualPage.hashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.hashResponse, emptyResponseHash)
	}
}
