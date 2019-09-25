package worker

import (
	"testing"

	"github.com/rishubhjain/web-crawler/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitialize(t *testing.T) {
	workerPool := WorkerPool{
		MaxWorkers: 1,
		Fn:         func(work *Work) {},
	}
	workerPool.Initialize()
}

func TestAddWork(t *testing.T) {
	workerPool := WorkerPool{
		MaxWorkers: 1,
		Fn:         func(work *Work) {},
	}
	workerPool.Initialize()

	mockWorker := new(tests.WorkerMock)
	mockWorker.On("Start", mock.Anything).Return(nil)

	work := Work{Depth: 1}
	workerPool.AddWork(&work)
	assert.Equal(t, len(workerPool.workers), 1)
}
