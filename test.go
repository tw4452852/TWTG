package main

import (
	"fmt"
	"strings"
)

func main() {
	mapFunc := func (r rune) rune {
		if r > 255 {
			println("!")
			return -r
		}
		return r
	}

	fmt.Print(strings.Map(mapFunc, "Hello world ??\n"))
}
