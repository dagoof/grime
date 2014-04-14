package grime

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	D := time.Millisecond * 10
	ticker := Tick(NewFixedStepper(D))

	a := <-ticker
	b := <-ticker
	c := <-ticker

	if b.Sub(a) != D {
		t.Errorf("one tick should take %s duration\n", D)
	}

	if c.Sub(a) != D*2 {
		t.Error("start to end should be two durations long")
	}

}
