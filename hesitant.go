// Package grime extends package time with some useful variations of
// standard types.
package grime

import (
	"time"
)

// HesitantTicker holds a channel that delivers ticks along an interval, but
// also aligned along that interval. The behaviour is such that a
// HesitantTicker of duration 5s will deliver every 5s starting 
// at 0s into every minute - continuing at 5s, 10s, 15s, and so on.
type HesitantTicker struct {
	C    chan time.Time
	ticker *time.Ticker
	stop chan struct{}
}

// Turn off the ticker. Does not close the channel, but prevents further ticks
// from being sent.
func (t *HesitantTicker) Stop() {
	close(t.stop)
}

// NewHesitantTicker returns a new HesitantTicker containing a channel that will
// send the time with a period specified by the duration argument. Ticks are
// dropped if a slow receiver is unable to keep up with the rate of ticks.
func NewHesitantTicker(d time.Duration) *HesitantTicker {
	ticker := &HesitantTicker{
		nil,
		make(chan time.Time, 1),
		make(chan struct{}),
	}

	go func(t *HesitantTicker) {
		var last time.Time
		time.Sleep(time.Now().Add(d).Truncate(d).Sub(time.Now()))

		t.ticker = time.NewTicker(d)
		for {
			last = <-t.Ticker.C

			select {
			case <-t.stop:
				t.Ticker.Stop()
				return
			case t.C <- last:
			default:
			}
		}
	}(ticker)

	return ticker
}

// HesitantTick is a convenience wrapper for NewHesitantTicker providing access
// to the ticking channel only. Useful for clients that have no need to shut
// down the channel.
// Analagous to time.Tick.
func HesitantTick(d time.Duration) <-chan time.Time {
	return NewHesitantTicker(d).C
}
