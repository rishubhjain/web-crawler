package tests

import (
	"context"
	"net/url"

	"github.com/rishubhjain/web-crawler/types"

	"github.com/stretchr/testify/mock"
)

// Mocking Parser Interface
type HTTPParserMock struct {
	mock.Mock
}

func (h *HTTPParserMock) Parse(ctx context.Context, site *types.Site) (err error) {
	newSite := types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	},
		Links: nil}
	site.Links = append(site.Links, &newSite)
	return nil
}

// Mocking Worker Interface
type WorkerMock struct {
	mock.Mock
}

func (w *WorkerMock) Start() {
	return
}
