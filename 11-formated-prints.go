package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guessed data type
	// %T : Type of supplied value

	fmt.Printf(
		"%d %c %f %t %s %o %x %v %T\n",
		1, 'A', 3.14, true, "stuff", 1, 1, 1, 1,
	)

	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	p1(sp1)
}
