package worker

import (
	"sync"

	"github.com/rishubhjain/web-crawler/fetch"
	"github.com/rishubhjain/web-crawler/types"
)

// Worker abstracts functionality of worker
type Worker interface {
	Start()
}

type worker struct {
	Fn    func(*Work)
	Queue chan Work
}

// Work struct stores arguments of function to be called
type Work struct {
	Site    *types.Site
	Depth   int
	Fetcher fetch.Fetcher
	Visited *types.Set
	Wg      *sync.WaitGroup
}

// Start a worker
func (w *worker) Start() {
	go w.run()
}

func (w *worker) run() {
	for work := range w.Queue {
		w.Fn(&work)
	}
}
