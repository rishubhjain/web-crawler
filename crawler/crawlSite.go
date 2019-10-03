package crawler

import (
	"context"
	"net/http"
	"time"

	cerror "github.com/rishubhjain/web-crawler/errors"
	HTTPParser "github.com/rishubhjain/web-crawler/parse"
	"github.com/rishubhjain/web-crawler/worker"

	log "github.com/sirupsen/logrus"
)

// CrawlSite abtracts crawling single site functionality
type CrawlSite interface {
	Run(*worker.Work)
}

type crawlSite struct {
	workerPool worker.WorkerPool
	parser     HTTPParser.Parser
	jobQueue   chan worker.Work
}

// NewCrawlSite returns a CrawlSite instance
func NewCrawlSite() CrawlSite {
	// ToDo: Make this configurable
	return &crawlSite{
		workerPool: worker.WorkerPool{
			MaxWorkers: 10000,
		},
		// Using default http client for now
		parser:   HTTPParser.NewHTTPParser(http.DefaultClient),
		jobQueue: make(chan worker.Work, 10000),
	}
}

// Run function crawls through each URL and adds sites to sitemap tree
func (w *crawlSite) Run(work *worker.Work) {

	w.workerPool.Fn = w.run

	// Initialize the Worker Pool
	w.workerPool.Initialize()

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
		}
	}
}

func (w *crawlSite) run(work *worker.Work) {
	// Check whether the URL has been already visited or not
	if work.Visited.Has(work.Site.URL.String()) {
		return
	}

	// Check whether max depth has reached or not
	if work.Depth <= 0 {
		return
	}

	work.Visited = work.Visited.Add(work.Site.URL.String())

	err := w.parser.Parse(context.Background(), work.Site)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": work.Site.URL.String()}).Error(cerror.ErrFetchFailed)
		return
	}

	// Loop through all the URLs to crawl through each url
	for _, childURL := range work.Site.Links {
		job := worker.Work{
			Site:    childURL,
			Depth:   work.Depth - 1,
			Visited: work.Visited,
		}
		w.workerPool.AddWork(&job)
	}
}
