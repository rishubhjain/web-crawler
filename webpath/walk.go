package webpath

import (
	"context"
	"net/http"

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
	// go w.walk(work)
	// faninApproach(w)

	w.workerPool.AddWork(work)
	w.workerPool.Wait()
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
			"URL": work.Site.URL.String()}).Error(cerror.ErrURLfetchFailed)
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

// func faninApproach(w *walkURL) {
// 	loop := true
// 	for loop {
// 		select {
// 		case doWork := <-w.jobQueue:
// 			w.workerPool.AddWork(&doWork)
// 		// ToDo: Make this configurable
// 		// This will kill the process if the time exceeds
// 		// a certain limit. this should be a kind of time out
// 		case <-time.After(time.Hour):
// 			loop = false
// 			return
// 		// To check whether the all the jobs are finished or not
// 		// This has an edge case where temperorily there is no
// 		// job in the queue
// 		case <-time.Tick(10 * time.Second):
// 			if w.workerPool.JobCount() == 0 {
// 				loop = false
// 			}
// 		}
// 	}
// 	return
// }
