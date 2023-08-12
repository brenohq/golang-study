package main

import (
	"fmt"
)

var p1 = fmt.Println

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	p1("4! =", factorial(4))
}
