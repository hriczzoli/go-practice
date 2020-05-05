package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	//This will pull the value off of the channel
	//will print 42 true
	v, ok := <-c
	fmt.Println(v, ok)

	//Closing the channel
	close(c)

	//This cannot pull the value off of the channel
	//will print 0 false
	v, ok = <-c
	fmt.Println(v, ok)
}
