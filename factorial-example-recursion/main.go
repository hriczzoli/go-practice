package main

import "fmt"

func main() {
	fmt.Println(factorial(7))
}

func factorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
