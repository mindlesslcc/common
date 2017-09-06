// package multichan
// implements a chan offer one put multi output
package multichan

// MultiChan contains multi chans
type MultiChan struct {
	chans []chan interface{}
}

// NewMultiChan new a multi chan
func NewMultiChan() (*MultiChan, error) {
	return &MultiChan{
		chans: make([]chan interface{}, 0, 1),
	}, nil
}

// Signal will send msg to all chans
func (mchan *MultiChan) Signal(msg interface{}) error {
	for _, ch := range mchan.chans {
		ch <- msg
	}
	return nil
}

// GetChan return a channel for user to recv
func (mchan *MultiChan) GetChan() (chan interface{}, error) {
	ch := make(chan interface{})
	mchan.chans = append(mchan.chans, ch)
	return ch, nil
}
