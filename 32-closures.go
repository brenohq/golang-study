package main

import (
	"fmt"
)

var p1 = fmt.Println

func useFunc(f func(int, int) int, x, y int) {
	p1("Answer:", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	// simple closue
	intSum := func(x, y int) int { return x + y }
	p1("5 + 3 =", intSum(5, 3))

	// changing values defined outside the closure
	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}
	changeVar()
	p1("samp1:", samp1)

	// passing function as argument
	useFunc(sumVals, 5, 8)
}
