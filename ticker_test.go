package grime

import (
	"fmt"
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

func ExampleFixedStepper() {
	var (
		D   = time.Millisecond * 10
		now = time.Now().Truncate(D)
	)

	stepper := NewFixedStepper(D)
	stepper.Start = now

	ticker := Tick(stepper)

	fmt.Println((<-ticker).Sub(now))
	fmt.Println((<-ticker).Sub(now))
	fmt.Println((<-ticker).Sub(now))

	// Output:
	// 10ms
	// 20ms
	// 30ms
}

func ExampleFixedStepper_startPast() {
	start, err := time.Parse(time.StampMilli, "Apr 14 15:49:03.000")
	if err != nil {
		fmt.Println(err)
	}

	D := time.Millisecond * 100
	stepper := NewFixedStepper(D)
	stepper.Start = start
	stepper.Grace = D / 4

	ticker := Tick(stepper)

	fmt.Println((<-ticker).Format(time.StampMilli))
	fmt.Println((<-ticker).Format(time.StampMilli))
	fmt.Println((<-ticker).Format(time.StampMilli))
	// Output:
	// Apr 14 15:49:03.100
	// Apr 14 15:49:03.200
	// Apr 14 15:49:03.300
}
