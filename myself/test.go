package main

import (
	"fmt"
	"./greetings"
)

func main() {
	var a int64
	fmt.Printf("Size of int is %d\n", unsafe.Sizeof(a))
}
