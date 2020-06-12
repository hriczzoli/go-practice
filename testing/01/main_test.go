package main

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	got := Foo(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
