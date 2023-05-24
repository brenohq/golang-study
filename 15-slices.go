package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// creating slice
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	p1("Slice size:", len(sl1))

	// iterating through
	for i := 0; i < len(sl1); i++ {
		p1(sl1[i])
	}

	for _, x := range sl1 {
		p1(x)
	}

	// creating array
	sArr := [5]int{1, 2, 3, 4, 5}

	// creating slice from an existent array
	sl3 := sArr[0:2]

	p1("First three:", sArr[:3])
	p1("Last three:", sArr[2:])

	// operations in slices that affect arrays
	sArr[0] = 10
	p1("sl3:", sl3)

	sl3[0] = 1
	p1("sArr:", sArr)

	sl3 = append(sl3, 12)
	p1("sl3:", sl3)
	p1("sArr:", sArr)

	sl4 := make([]string, 6)
	p1("sl4:", sl4)
	p1("sl4[0]:", sl4[0])
}
