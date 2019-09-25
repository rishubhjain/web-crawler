package webpath

import (
	"errors"
	"net/url"
	"sync"
	"testing"

	"github.com/rishubhjain/web-crawler/tests"
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/worker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWalk(t *testing.T) {

	mockFetcher := new(tests.HTTPFetcherMock)
	mockFetcher.On("Fetch", mock.Anything).Return(nil)

	site := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	depth := 1

	// Create client from the mocked fetcher
	client := tests.HTTPFetcherMock{}

	visited := &types.Set{}
	var wg sync.WaitGroup

	mockWorker := new(tests.WorkerMock)
	mockWorker.On("Start", mock.Anything).Return(nil)

	wg.Add(1)
	work := worker.Work{
		Site:  &site,
		Depth: depth,
		// Using default http client for now
		Fetcher: &client,
		Visited: visited,
		Wg:      &wg,
	}

	walkObj := NewWalkURL()
	walkObj.Walk(&work)
	//wg.Wait()
	assert.Equal(t, len(site.Links), 1)

	site.Links = nil
	wg.Add(1)
	walkObj.Walk(&work)

	assert.Equal(t, len(site.Links), 0)

	mockFetcher.On("Fetch", mock.Anything).Return(errors.New("Test Error"))
	wg.Add(1)
	walkObj.Walk(&work)

	assert.Equal(t, len(site.Links), 0)

}
