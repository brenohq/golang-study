package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	f4 := 10
	var f4Ptr *int = &f4

	p1("f4 address:", f4Ptr)
	p1("f4 value:", *f4Ptr)

	*f4Ptr = 12
	p1("f4 value after re assignment:", *f4Ptr)
}
