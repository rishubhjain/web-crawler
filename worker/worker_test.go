package worker

import (
	"testing"

	"github.com/rishubhjain/web-crawler/types"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	poolObj := WorkerPool{
		MaxWorkers: 1,
		Fn:         func(work *Work) { work.Depth = 2 },
	}

	poolObj.Initialize()

	site := &types.Site{URL: nil}
	work := Work{
		Site:  site,
		Depth: 1,
		// the visited URl is initiated here, but it can be
		// fetched from a central store and passed on.
		Visited: &types.Set{},
	}
	poolObj.AddWork(&work)
	assert.Equal(t, poolObj.JobCount(), 1)
}
