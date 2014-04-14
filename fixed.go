package grime

import (
	"time"
)

// FixedStepper delivers steps at an interval. Can potentially start in the past
// or future based on an optional Start offset.
type FixedStepper struct {
	D     time.Duration
	Start time.Time
	tick  time.Time
}

// Step forward according to duration.
func (t *FixedStepper) Step() time.Time {
	if t.tick.IsZero() {
		if t.Start.IsZero() {
			t.tick = time.Now().Add(-t.D)
		} else {
			t.tick = t.Start
		}
	}

	t.tick = t.tick.Add(t.D)
	time.Sleep(t.tick.Sub(time.Now()))
	return t.tick
}

// NewFixedStepper creates a new FixedStepper with no start time specified.
func NewFixedStepper(D time.Duration) *FixedStepper {
	return &FixedStepper{D: D}
}
