package main

import "fmt"

var p1 = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	p1("5 + 4 = ", getSumGen(5, 4))
	p1("5.4 + 4.7 = ", getSumGen(4.3, 4.7))
}
