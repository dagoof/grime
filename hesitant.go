package grime

import (
	"time"
)

// NewHesitantStepper delivers steps at an interval, but
// also aligned along that interval. The behaviour is such that a
// HesitantStepper of duration 5s will deliver every 5s starting
// at 0s into every minute - continuing at 5s, 10s, 15s, and so on.
func NewHesitantStepper(D time.Duration) *FixedStepper {
	return &FixedStepper{D, time.Now().Truncate(D)}
}
