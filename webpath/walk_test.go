package webpath

import (
	"errors"
	"net/url"
	"testing"

	"github.com/rishubhjain/web-crawler/tests"
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/worker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRun(t *testing.T) {

	mockFetcher := new(tests.HTTPFetcherMock)
	mockFetcher.On("Fetch", mock.Anything).Return(nil)

	site := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	depth := 1

	visited := &types.Set{}

	mockWorker := new(tests.WorkerMock)
	mockWorker.On("Start", mock.Anything).Return(nil)

	work := worker.Work{
		Site:    &site,
		Depth:   depth,
		Visited: visited,
	}

	crawlObj := New()
	crawlObj.Run(&work)
	assert.Equal(t, len(site.Links), 17)

	site.Links = nil
	crawlObj.Run(&work)

	assert.Equal(t, len(site.Links), 0)

	mockFetcher.On("Fetch", mock.Anything).Return(errors.New("Test Error"))

	crawlObj.Run(&work)

	assert.Equal(t, len(site.Links), 0)

}
