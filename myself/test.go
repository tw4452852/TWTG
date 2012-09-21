package main

import (
	"fmt"
	"./stack"
)

func main() {
	s := stack.NewStack()

	s.Push(1)
	s.Push("hello tw")
	fmt.Println(s)
}
