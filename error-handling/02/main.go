package main

import (
	"fmt"
)

type customErr struct {
	message string
}

func (c customErr) Error() string {
	return fmt.Sprintf("the following error occured in the program: %v", c.message)
}

func main() {
	cErr := customErr{"this message should be specific for the error where it happened"}
	foo(cErr)
}

func foo(err error) {
	fmt.Println(err)

	//assertion
	fmt.Println(err.(customErr).message)
}
