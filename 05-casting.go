package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var p1 = fmt.Println

func main() {
	// casting float to int
	cV1 := 1.5
	cV2 := int(cV1)
	p1(cV2)

	// casting string to int
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	p1(cV4, err, reflect.TypeOf(cV4))

	// casting int to string
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	p1(cV6, reflect.TypeOf(cV6))

	// casting string to float
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		p1(cV8, reflect.TypeOf(cV8))
	}
}
