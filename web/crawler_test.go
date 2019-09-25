package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/tests"
	"github.com/rishubhjain/web-crawler/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking Fetcher Interface
type httpFetcherMock struct {
	mock.Mock
}

// NewHTTPFetcherMock returns a crawler instance
func NewHTTPFetcherMock(client *http.Client) fetch.Fetcher {
	return &httpFetcherMock{}
}

func (h *httpFetcherMock) Fetch(ctx context.Context,
	site *types.Site) (err error) {
	newSite := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	site.Links = append(site.Links, &newSite)
	return nil
}

func TestCrawl(t *testing.T) {

	body := tests.ParseHTML(t, "../tests/fixtures/test.html")
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Send response to be tested
		rw.Write([]byte(body))
	}))
	// Close the server when test finishes
	defer server.Close()

	mockFetcher := new(httpFetcherMock)
	mockFetcher.On("Fetch", mock.Anything).Return(nil)
	crawler := NewCrawler()
	baseURL := server.URL
	depth := 1
	site, _ := crawler.Crawl(baseURL, depth)
	assert.Equal(t, len(site.Links), 2)

	baseURL = "h//*goo%&^gle"
	_, err := crawler.Crawl(baseURL, depth)
	assert.NotNil(t, err)

}
