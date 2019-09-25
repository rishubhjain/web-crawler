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
	Crawl(hostURL string, depth int) (*types.Site, error)
}

// Structure to implement Crawler interface
type crawler struct{}

// NewCrawler returns Crawler instance
func NewCrawler() Crawler {
	return &crawler{}
}

// Crawl function parses the host URL and crawls through all the URL's in
// the site page recursively
func (c *crawler) Crawl(hostURL string, depth int) (*types.Site, error) {

	// Parse the site URL
	URL, err := utils.Parse(hostURL)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": hostURL}).Error("Failed to parse URL")
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
