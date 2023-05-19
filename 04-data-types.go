package main

import (
	"fmt"
	"reflect"
)

var p1 = fmt.Println

func main() {
	// int, float64, bool, string, rune
	// default vales are, 0, 0.0, false, "" and ""

	p1(reflect.TypeOf(25))
	p1(reflect.TypeOf(3.14))
	p1(reflect.TypeOf(true))
	p1(reflect.TypeOf("Hello"))
	p1(reflect.TypeOf('⚽'))
}
