package bufferwaitgroup

import (
	"sync"
	"sync/atomic"
)

const (
	RUN_BUF_SIZE = 1024
)

type runner func()

type BufferWaitGroup struct {
	sync.WaitGroup
	size      int32
	capacity  int32
	runnerBuf chan runner
}

func NewBufferWaitGroup(size int32) *BufferWaitGroup {
	return &BufferWaitGroup{
		capacity:  size,
		runnerBuf: make(chan runner, RUN_BUF_SIZE),
	}
}

func (bwg *BufferWaitGroup) isFull() bool {
	return atomic.LoadInt32(&bwg.size) >= bwg.capacity
}

func (bwg *BufferWaitGroup) add(num int32) {
	atomic.AddInt32(&bwg.size, num)
	bwg.Add(int(num))
}

func (bwg *BufferWaitGroup) done() {
	atomic.AddInt32(&bwg.size, -1)
	bwg.Done()
}

func (bwg *BufferWaitGroup) Wrap(f runner) {
	bwg.runnerBuf <- f
	if !bwg.isFull() {
		bwg.add(1)
		go bwg.work()
	}
}

func (bwg *BufferWaitGroup) work() {
	for f := range bwg.runnerBuf {
		f()
		if len(bwg.runnerBuf) == 0 {
			break
		}
	}
	bwg.done()
}
