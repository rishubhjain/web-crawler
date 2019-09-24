package web

import (
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"
	"github.com/rishubhjain/web-crawler/webpath"

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
	webpath.Walk(site, depth)
	return site, nil
}
