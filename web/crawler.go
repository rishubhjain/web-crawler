package web

import (
	"net/http"

	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"
	"github.com/rishubhjain/web-crawler/webpath"
	"github.com/rishubhjain/web-crawler/worker"

	log "github.com/sirupsen/logrus"
)

// Crawler abstracts crawling functionality
type Crawler interface {
	Crawl(baseURL string, depth int) (*types.Site, error)
}

// Structure to implement Crawler interface
type crawler struct {
}

// NewCrawler returns a crawler instance
func NewCrawler() Crawler {
	return &crawler{}
}

// Crawl function crawls through all URLs returns the base site
func (c *crawler) Crawl(baseURL string, depth int) (*types.Site, error) {

	// Parse the site URL
	URL, err := utils.Parse(baseURL)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": baseURL}).Error("Failed to parse URL")
		return nil, err
	}

	// Start crawling from the base site
	site := &types.Site{URL: URL}
	walkObj := webpath.NewWalkURL()

	// Creating work for the Walk function
	work := worker.Work{
		Site:  site,
		Depth: depth,
		// Using default http client for now
		Fetcher: fetch.NewHTTPFetcher(http.DefaultClient),
		Visited: &types.Set{},
		Wg:      nil,
	}

	walkObj.Walk(&work)
	return site, nil
}
