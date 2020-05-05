package main

import (
	"fmt"
)

//A program that [uts a 100 numbers on to a channel
//then pulls them off

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println("Pulled from channel: ", v)
	}

	fmt.Println("About to exit")
}
