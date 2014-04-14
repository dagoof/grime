package grime

import (
	"time"
)

// HesitantStepper delivers steps at an interval, but
// also aligned along that interval. The behaviour is such that a
// HesitantStepper of duration 5s will deliver every 5s starting
// at 0s into every minute - continuing at 5s, 10s, 15s, and so on.
type HesitantStepper struct {
	D     time.Duration
	Start time.Time
	tick  time.Time
}

// Step forward according to the initialzed duration.
func (t *HesitantStepper) Step() time.Time {
	if t.tick.IsZero() {
		if t.Start.IsZero() {
			t.tick = time.Now().Truncate(t.D)
		} else {
			t.tick = t.Start.Truncate(t.D)
		}
	}

	t.tick = t.tick.Add(t.D)
	time.Sleep(t.tick.Sub(time.Now()))
	return t.tick
}

// NewHesitantStepper creates a new HesitantStepper with no start time
// specified.
func NewHesitantStepper(D time.Duration) *HesitantStepper {
	return &HesitantStepper{D: D}
}
