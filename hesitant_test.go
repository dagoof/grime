package grime

import (
	"fmt"
	"time"
)

func ExampleHesitantStepper() {
	D := time.Millisecond * 50

	now := time.Now().Truncate(D)
	t := NewTicker(NewHesitantStepper(D))
	defer t.Stop()

	fmt.Println((<-t.C).Sub(now))
	fmt.Println((<-t.C).Sub(now))
	fmt.Println((<-t.C).Sub(now))
	// Output:
	// 50ms
	// 100ms
	// 150ms
}

func ExampleHesitantTicker_slowReceiver() {
	D := time.Millisecond * 100

	now := time.Now().Truncate(D)
	t := NewTicker(NewHesitantStepper(D))
	defer t.Stop()

	fmt.Println((<-t.C).Sub(now))
	time.Sleep(D * 2)

	fmt.Println((<-t.C).Sub(now))
	fmt.Println((<-t.C).Sub(now))
	// Output:
	// 100ms
	// 200ms
	// 400ms
}
