package worker

import (
	"testing"

	"github.com/rishubhjain/web-crawler/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initialize() *WorkerPool {
	workerPool := WorkerPool{
		MaxWorkers: 1,
		Fn:         func(work *Work) {},
	}
	workerPool.Initialize()
	return &workerPool
}
func TestInitialize(t *testing.T) {
	workerPool := initialize()
	assert.Equal(t, len(workerPool.jobs), 0)
}

func TestAddWork(t *testing.T) {
	workerPool := initialize()

	mockWorker := new(tests.WorkerMock)
	mockWorker.On("Start", mock.Anything).Return(nil)

	work := Work{Depth: 1}
	workerPool.AddWork(&work)
	assert.Equal(t, len(workerPool.workers), 1)
}
