package main

import (
	"fmt"
)

func main() {
	//With anonymus func literal launching a Goroutine
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	//With buffer channel
	b := make(chan int, 1)
	b <- 42

	fmt.Println(<-b)

	//If using a channel without a goroutine or buffer channel we end up with a deadlock
	// - Error: all Goroutines are asleep
}
