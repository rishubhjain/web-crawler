package worker

import "testing"

func TestStart(t *testing.T) {
	workChan := make(chan Work)
	workerObj := worker{
		Fn:    func(work *Work) {},
		Queue: workChan,
	}
	workerObj.Start()
}
