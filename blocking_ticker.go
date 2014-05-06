package grime

import (
	"time"
)

// BlockingTicker holds a channel that delivers ticks according to a
// Stepper. All steps are read and sent through the contained channel.
type BlockingTicker struct {
	Stepper
	C    chan time.Time
	stop chan struct{}
}

// Stop the ticker. No more ticks will be sent after stop has been called, but
// it does not close the tick channel.
func (t *BlockingTicker) Stop() {
	close(t.stop)
}

// NewBlockingTicker creates a Ticker containing a channel that will send the
// ticks according to a Stepper. Steps will never be discarded.
func NewBlockingTicker(s Stepper) *BlockingTicker {
	ticker := &BlockingTicker{
		s,
		make(chan time.Time),
		make(chan struct{}),
	}

	go func(t *BlockingTicker) {
		tick := t.Step()

		for {
			select {
			case <-t.stop:
				return
			case t.C <- tick:
			}

			tick = t.Step()

		}
	}(ticker)

	return ticker
}

// BlockingTick is a convenience wrapper to create a new BlockingTicker
// when the client has no need to shut down the ticker.
func BlockingTick(s Stepper) <-chan time.Time {
	return NewBlockingTicker(s).C
}
