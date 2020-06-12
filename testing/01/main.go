//A file that contains a simple function which we can test
//in the corresponding test file
package main

func main() {
	Foo(12)
}

//Foo will be used in a test that's why it needs to be exported
func Foo(i int64) int64 {
	return i
}
