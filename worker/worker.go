package worker

import (
	"sync"

	"github.com/rishubhjain/web-crawler/types"
)

// Worker abstracts functionality of worker
type Worker interface {
	Start()
}

type worker struct {
	Fn    func(*Work)
	Queue chan Work
	Wg    *sync.WaitGroup
	End   chan bool
}

// Work struct stores arguments of function to be called
// ToDo: also store errors. and move Work to type package
type Work struct {
	Site    *types.Site
	Depth   int
	Visited *types.Set
}

// Start a worker
func (w *worker) Start() {
	go w.run()
}

func (w *worker) run() {
	for {
		// make the channel available
		select {
		case work := <-w.Queue:
			w.Fn(&work)
			w.Wg.Done()
		// To stop the worker
		case <-w.End:
			return
		}
	}
}

// Stop function is used to stop the worker
func (w *worker) Stop() {
	w.End <- true
}
