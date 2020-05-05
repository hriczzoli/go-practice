package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Creating custom type to be saved as the value on to the channel
type cmap map[int]string

func main() {
	var wg sync.WaitGroup
	ch := make(chan cmap)
	var result []cmap

	go func() {
		//Launching 10 GOroutines - each one puts 10 numbers on the channel
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					str := "Goroutine #" + strconv.Itoa(m)
					mi := cmap{i*10 + m: str}
					ch <- mi
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(ch)
	}()

	//Looping over the channel and pulling the values off of it
	//while appending the values to a slice of our custom type
	for i := 0; i < 100; i++ {
		result = append(result, <-ch)
	}

	formatResult(result)

	fmt.Println("About to exit")
}

func formatResult(r []cmap) {
	for i, v := range r {
		for k, s := range v {
			fmt.Printf("At index %v\thave the number:%v\t -coming from:%v\n", i, k, s)
		}
	}
}
