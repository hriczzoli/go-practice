package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	//Simple error: invalid operation: <-cs (receive from send-only type chan<- int)
	// fixing it by making the channel bi-directional
}
