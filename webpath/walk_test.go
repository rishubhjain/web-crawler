package webpath

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"testing"

	"github.com/rishubhjain/web-crawler/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking Fetcher Interface
type HTTPFetcherMock struct {
	mock.Mock
}

func (h *HTTPFetcherMock) Fetch(ctx context.Context, site *types.Site) (err error) {
	newSite := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	site.Links = append(site.Links, &newSite)
	return nil
}

func TestWalk(t *testing.T) {

	mockFetcher := new(HTTPFetcherMock)
	mockFetcher.On("Fetch", mock.Anything).Return(nil)
	site := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	depth := 1

	// Create client from the mocked fetcher
	client := HTTPFetcherMock{}

	visited := &types.Set{}
	var wg sync.WaitGroup

	wg.Add(1)
	walk(&site, depth, &client, visited, &wg)
	wg.Wait()
	assert.Equal(t, len(site.Links), 1)

	site.Links = nil
	wg.Add(1)
	walk(&site, 0, &client, visited, &wg)
	wg.Wait()
	assert.Equal(t, len(site.Links), 0)

	Walk(&site, 0)
	assert.Equal(t, len(site.Links), 0)

	mockFetcher.On("Fetch", mock.Anything).Return(errors.New("Test Error"))
	wg.Add(1)
	walk(&site, 0, &client, visited, &wg)
	wg.Wait()
	assert.Equal(t, len(site.Links), 0)

}
