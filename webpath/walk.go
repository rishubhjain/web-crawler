package webpath

import (
	"context"
	"sync"

	"github.com/rishubhjain/web-crawler/worker"

	log "github.com/sirupsen/logrus"
)

// WalkURL abtracts walk functionality
type WalkURL interface {
	Walk(*worker.Work)
}

type walkURL struct {
	workerPool worker.WorkerPool
}

// NewWalkURL returns a walkURL instance
func NewWalkURL() WalkURL {
	return &walkURL{
		workerPool: worker.WorkerPool{
			MaxWorkers: 1000, // Make this configurable
		},
	}
}

// Walk walks through each URL and creates the tree
func (w *walkURL) Walk(work *worker.Work) {
	var wg sync.WaitGroup

	w.workerPool.Fn = w.walk

	// Initialize the Worker Pool
	w.workerPool.Initialize()
	work.Wg = &wg

	wg.Add(1)

	w.workerPool.AddWork(work)
	wg.Wait()
}

func (w *walkURL) walk(work *worker.Work) {
	defer work.Wg.Done()
	// Check whether the URL has been already visited or not
	if work.Visited.Has(work.Site.URL.String()) {
		return
	}

	// Check whether max depth has reached or not
	if work.Depth <= 0 {
		return
	}

	work.Visited = work.Visited.Add(work.Site.URL.String())

	// Fetch all URLs in the site
	err := work.Fetcher.Fetch(context.Background(), work.Site)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": work.Site.URL.String()}).Error("Failed to fetch urls")
		return
	}

	// Loop through all the URLs to walk through each url
	for _, childURL := range work.Site.Links {
		work.Wg.Add(1)
		work := worker.Work{
			Site:    childURL,
			Depth:   work.Depth - 1,
			Fetcher: work.Fetcher,
			Visited: work.Visited,
			Wg:      work.Wg,
		}
		w.workerPool.AddWork(&work)
	}
}
