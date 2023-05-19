package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// enforcing type
	var vName string = "Breno"
	p1(vName)

	// inferring type
	var v1, v2 = 1.2, 3.4
	p1(v1, v2)

	// reassigning value
	v3 := 2.4
	v3 = 5.4
	p1(v3)
}
