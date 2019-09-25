package worker

// WorkerPool is a wrapper to manage a set of Workers efficiently
type WorkerPool struct {
	MaxWorkers int
	Fn         func(*Work)
	workers    []worker
	jobs       chan Work
}

// Initialize the WorkerPool
func (pool *WorkerPool) Initialize() {
	pool.jobs = make(chan Work, pool.MaxWorkers)
}

// AddWork adds work to the WorkerPool
func (pool *WorkerPool) AddWork(work *Work) {
	if len(pool.workers) < pool.MaxWorkers {
		worker := worker{
			Queue: pool.jobs,
			Fn:    pool.Fn,
		}
		worker.Start()
		pool.workers = append(pool.workers, worker)
	}
	pool.jobs <- *work
}
