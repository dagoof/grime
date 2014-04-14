// Package grime extends package time with some useful variations of
// standard types.
package grime

import (
	"time"
)

// Stepper is any kind of type that can deliver time steps, or ticks.
type Stepper interface {
	Step() time.Time
}

// Ticker holds a channel that delivers ticks according to a Stepper. All steps
// are read, but if slow recievers fail to keep up, steps may be discarded.
type Ticker struct {
	Stepper
	C    chan time.Time
	stop chan struct{}
}

// Stop the ticker. No more ticks will be sent after stop has been called, but
// it does not close the tick channel.
func (t *Ticker) Stop() {
	close(t.stop)
}

// NewTicker creates a Ticker containing a channel that will send the ticks
// according to a Stepper. Steps may be discarded if a receiver is unable to
// keep up.
func NewTicker(s Stepper) *Ticker {
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
			default:
			}

			tick = t.Step()

		}
	}(ticker)

	return ticker
}

// Tick is a convenience wrapper to create a new Ticker when the client has no
// need to shut down the ticker.
func Tick(s Stepper) <-chan time.Time {
	return NewTicker(s).C
}
