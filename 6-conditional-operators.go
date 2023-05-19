package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// conditional operators : > < >= <= == !=
	// logical operators && || !

	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		p1("Important birthday!")
	} else if (iAge == 21) || (iAge == 50) {
		p1("Also a important birthday!")
	} else if iAge >= 65 {
		p1("Getting old, so it is a important birthday too!")
	} else {
		p1("Whatever birthday!")
	}

	p1("!true =", !true)
}
