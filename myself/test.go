package main

import (
	"fmt"
	"./fibo"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, " is ", fibo.Fibonacci(i))
	}
}
