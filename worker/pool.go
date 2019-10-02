package worker

import "sync"

// WorkerPool is a wrapper to manage a set of Workers efficiently
type WorkerPool struct {
	MaxWorkers int
	Fn         func(*Work)
	workers    []worker
	jobs       chan Work
	Wg         sync.WaitGroup
}

// Initialize initializes the WorkerPool
func (pool *WorkerPool) Initialize() {
	pool.jobs = make(chan Work, pool.MaxWorkers)
}

// AddWork adds work to the WorkerPool
func (pool *WorkerPool) AddWork(work *Work) {
	var workerObj worker
	if len(pool.workers) < pool.MaxWorkers {
		workerObj = worker{
			Queue: pool.jobs,
			Fn:    pool.Fn,
			Wg:    &pool.Wg,
			End:   make(chan bool),
		}
		workerObj.Start()
		pool.workers = append(pool.workers, workerObj)
	}
	pool.Wg.Add(1)
	// This allocates data to goroutine randomly
	// hence unequal distribution.
	pool.jobs <- *work
}

// Wait function waits for all the workers to finish
func (pool *WorkerPool) Wait() {
	pool.Wg.Wait()
}

// StopWorkers will stop all workers
func (pool *WorkerPool) StopWorkers() {
	for i, worker := range pool.workers {
		worker.Stop()
		pool.workers = append(pool.workers[:i], pool.workers[i+1:]...)
	}
}

// JobCount returns the amount of work pending in the queue
func (pool *WorkerPool) JobCount() int {
	return len(pool.jobs)
}
