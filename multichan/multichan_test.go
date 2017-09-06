package multichan

import (
	"testing"
	"time"
)

func Test_MultiChan(t *testing.T) {
	var err error
	result := 5
	mc, err := NewMultiChan()
	if err != nil {
		t.Errorf("error when new multi-channel : %s\n", err)
		t.FailNow()
	}

	var ch chan interface{}
	chans := make([]chan interface{}, 0, 5)
	for i := 0; i < 5; i++ {
		ch, err = mc.GetChan()
		if err != nil {
			t.Errorf("error when new multi-channel : %s\n", err)
			t.FailNow()
		}
		chans = append(chans, ch)
		go func(i int) {
			<-chans[i]
			result--
		}(i)
	}

	mc.Signal("stop")

	time.Sleep(time.Second * 5)
	if result != 0 {
		t.Errorf("result not equal 0!\n")
		t.FailNow()
	}
}
