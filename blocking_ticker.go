package grime

import (
	"time"
)

// NewBlockingTicker creates a Ticker containing a channel that will send the
// ticks according to a Stepper. Steps will never be discarded.
func NewBlockingTicker(s Stepper) *Ticker {
	ticker := &Ticker{
		s,
		make(chan time.Time, 1),
		make(chan struct{}),
	}

	go func(t *Ticker) {
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
