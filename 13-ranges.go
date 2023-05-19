package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	aNums := []int{1, 2, 3}

	for _, num := range aNums {
		p1(num)
	}
}
