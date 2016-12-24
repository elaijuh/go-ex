package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Example 1: context.WithTimeout return a copy of parent context
	// context's Done channel is closed when timeout
	ctx1, cancel1 := context.WithTimeout(context.Background(), 50*time.Millisecond)

	select {
	case <-time.After(60 * time.Millisecond):
		fmt.Println("wait for 60 millisecond")
	case <-ctx1.Done():
		fmt.Println("Example 1:", ctx1.Err())
	}
	cancel1()

	// Example 2: context.WithCancel return a copy of parent context
	// context's Done channel is closed when cancel is called
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func(c *context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("Example 2:", ctx2.Err())
		}
	}(&ctx2)

	cancel2()
	time.Sleep(50 * time.Millisecond)

	// Example 3: context.WithValue returns a copy of parent context with key value pair associated
	// string type key is not suggested due to namespace issue, define a custom type for context key
	type ctxKeyType string
	mo := ctxKeyType("mo")
	po := ctxKeyType("po")
	ctx3 := context.WithValue(context.Background(), mo, "m")
	fmt.Println("Example 3: ctx3 mo's value", ctx3.Value(mo))
	ctx4 := context.WithValue(ctx3, po, "f")
	fmt.Println("Example 3: ctx4 mo's value", ctx4.Value(mo))
	fmt.Println("Example 3: ctx4 po's value", ctx4.Value(po))

}
