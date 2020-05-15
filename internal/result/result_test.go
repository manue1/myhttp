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
		Http: &mocks.Client{},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {

			page := client.Get(tc.url)
			tc.fn(t, tc.expectedURL, page)
		})
	}
}

func testSuccessAlreadySanitizedURL(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.URL {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.URL, expectedURL)
	}

	expectedHash := mocks.GetMockMD5(expectedURL)
	if expectedHash != actualPage.HashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.HashResponse, expectedHash)
	}
}

func testSuccessUnsanitizedURL(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.URL {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.URL, expectedURL)
	}

	expectedHash := mocks.GetMockMD5(expectedURL)
	if expectedHash != actualPage.HashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.HashResponse, expectedHash)
	}
}

func testFailureEmptyBody(t *testing.T, expectedURL string, actualPage Page) {
	if expectedURL != actualPage.URL {
		t.Errorf("unexpected url: got %s want %s",
			actualPage.URL, expectedURL)
	}

	emptyResponseHash := "d41d8cd98f00b204e9800998ecf8427e"
	if emptyResponseHash != actualPage.HashResponse {
		t.Errorf("unexpected md5 hash: got %s want %s",
			actualPage.HashResponse, emptyResponseHash)
	}
}
