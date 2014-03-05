package grime

import (
	"fmt"
	"time"
)

func ExampleHesitantTicker() {
	d := time.Millisecond * 100

	t := NewHesitantTicker(d)
	now := time.Now().Add(d).Truncate(d)
	defer t.Stop()

	fmt.Println((<-t.C).Truncate(d).Sub(now))
	fmt.Println((<-t.C).Truncate(d).Sub(now))
	fmt.Println((<-t.C).Truncate(d).Sub(now))
	// Output:
	// 100ms
	// 200ms
	// 300ms
}
