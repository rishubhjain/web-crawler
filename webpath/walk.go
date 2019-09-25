package webpath

import (
	"context"
	"net/http"
	"sync"

	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/types"

	log "github.com/sirupsen/logrus"
)

// Walk walks through each URL and creates the tree
func Walk(site *types.Site, depth int) {
	// Using default http client for now
	var wg sync.WaitGroup
	wg.Add(1)
	go walk(site, depth, fetch.NewHTTPFetcher(http.DefaultClient), &types.Set{}, &wg)
	wg.Wait()
}

func walk(site *types.Site, depth int, fetcher fetch.Fetcher, visited *types.Set, wg *sync.WaitGroup) {
	defer wg.Done()
	// Check whether the URL has been already visited or not
	if visited.Has(site.URL.String()) {
		return
	}

	// Check whether max depth has reached or not
	if depth <= 0 {
		return
	}

	visited = visited.Add(site.URL.String())

	// Fetch all URLs in the site
	err := fetcher.Fetch(context.Background(), site)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": site.URL.String()}).Error("Failed to fetch urls")
		return
	}

	for _, childURL := range site.Links {
		wg.Add(1)
		// Go-routines to fetch urls for all childurls of site
		go func(childURL *types.Site, visited *types.Set) {
			walk(childURL, depth-1, fetcher, visited, wg)
		}(childURL, visited)
	}
}
