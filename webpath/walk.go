package webpath

import (
	"context"
	"net/http"

	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/types"

	log "github.com/sirupsen/logrus"
)

// Walk walks through each URL and creates the tree
func Walk(site *types.Site, depth int) {
	// Using default http client for now
	walk(site, depth, fetch.NewGoqueryFetcher(http.DefaultClient), &types.Set{})
}

func walk(site *types.Site, depth int, fetcher fetch.Fetcher, visited *types.Set) {
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

	done := make(chan struct{})
	for _, childURL := range site.Links {
		// Go-routines to fetch urls for all childurls of site
		go func(childURL *types.Site, visited *types.Set) {
			walk(childURL, depth-1, fetcher, visited)
			done <- struct{}{}
		}(childURL, visited)
	}

	// Waiting for all the go-routines to finish
	for range site.Links {
		<-done
	}
}
