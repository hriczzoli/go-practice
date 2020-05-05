package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//Launching 10 GOroutines - each one puts 10 numbers on the channel
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}

	//Looping over the channel and pulling the values off of it
	for j := 0; j < 100; j++ {
		fmt.Println(j, <-ch)
	}

	fmt.Println("About to exit")
}
