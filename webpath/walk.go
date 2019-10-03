package webpath

import (
	"context"
	"net/http"
	"time"

	cerror "github.com/rishubhjain/web-crawler/errors"
	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/worker"

	log "github.com/sirupsen/logrus"
)

// WalkURL abtracts walk functionality
type WalkURL interface {
	Walk(*worker.Work)
}

type walkURL struct {
	workerPool worker.WorkerPool
	fetcher    fetch.Fetcher
	jobQueue   chan worker.Work
}

// NewWalkURL returns a walkURL instance
func NewWalkURL() WalkURL {
	// ToDo: Make this configurable
	return &walkURL{
		workerPool: worker.WorkerPool{
			MaxWorkers: 10000,
		},
		// Using default http client for now
		fetcher:  fetch.NewHTTPFetcher(http.DefaultClient),
		jobQueue: make(chan worker.Work, 10000),
	}
}

// Walk function walks through each URL and adds site to tree
func (w *walkURL) Walk(work *worker.Work) {

	w.workerPool.Fn = w.walk

	// Initialize the Worker Pool
	w.workerPool.Initialize()

	// Fan In approach - ignore
	// {
	// 	go w.walk(work)
	// }

	stop := make(chan bool, 1)

	// Start the crawling process with head URL
	w.workerPool.AddWork(work)

	// Wait for all the goroutines to finish
	go func() {
		w.workerPool.Wait()
		stop <- true
	}()

	// To stop the crawler when either all goroutines finish the work
	// or time taken exceeds 1 hour
	for {
		select {
		case <-stop:
			return
		// TODO: Make this configurable
		case <-time.After(time.Hour):
			return
			// Fan In approach
			/* case doWork := <-w.jobQueue:
				w.workerPool.AddWork(&doWork)
			case <-time.Tick(10 * time.Second):
				if w.workerPool.JobCount() == 0 {
					return
				} */
		}
	}
}

func (w *walkURL) walk(work *worker.Work) {
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
	err := w.fetcher.Fetch(context.Background(), work.Site)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": work.Site.URL.String()}).Error(cerror.ErrFetchFailed)
		return
	}

	// Loop through all the URLs to walk through each url
	for _, childURL := range work.Site.Links {
		job := worker.Work{
			Site:    childURL,
			Depth:   work.Depth - 1,
			Visited: work.Visited,
		}
		//w.jobQueue <- job
		w.workerPool.AddWork(&job)
	}
}
