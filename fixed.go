package grime

import (
	"time"
)

// FixedStepper delivers steps at an interval. Can potentially start in the past
// or future based on an optional Start offset.
type FixedStepper struct {
	D    time.Duration
	tick time.Time
}

// Start sets the tick of the ticker if the stepper has not yet been started
func (s *FixedStepper) Start(t time.Time) {
	if s.tick.IsZero() {
		s.tick = t
	}
}

// Step forward according to duration.
func (s *FixedStepper) Step() time.Time {
	if s.tick.IsZero() {
		s.tick = time.Now().Add(-s.D)
	}

	s.tick = s.tick.Add(s.D)
	time.Sleep(s.tick.Sub(time.Now()))
	return s.tick
}

// NewFixedStepper creates a new FixedStepper with no start time specified.
func NewFixedStepper(D time.Duration) *FixedStepper {
	return &FixedStepper{D: D}
}
