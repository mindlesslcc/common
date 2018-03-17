package bufferwaitgroup

import (
	"sync/atomic"
	"testing"
	"time"
)

const (
	GOROUTINE_MAX_NUM  = 2
	GOROUTINE_TEST_NUM = 5
)

func Test_BufferWaitGroup(t *testing.T) {
	var val int32
	bwg := NewBufferWaitGroup(2)
	for i := 0; i < GOROUTINE_TEST_NUM; i++ {
		bwg.Wrap(func() {
			atomic.AddInt32(&val, 1)
			time.Sleep(time.Second * 1)
		})
	}
	time.Sleep(time.Second / 10)
	if atomic.LoadInt32(&val) != GOROUTINE_MAX_NUM {
		t.Fatal("not equal ", string(GOROUTINE_MAX_NUM), " but ", atomic.LoadInt32(&val))
	}
	time.Sleep(3 * time.Second)
	bwg.Wait()
	if atomic.LoadInt32(&val) != GOROUTINE_TEST_NUM {
		t.Fatal("not equal ", string(GOROUTINE_TEST_NUM), " but ", atomic.LoadInt32(&val))
	}
}
