package main

import (
	// importing mypackage
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

func main() {
	// using mypackage defined public constant variable
	fmt.Println("Hello", stuff.Name)

	// using mypackage public method
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)

	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	// using mypackage data type and getters
	date := stuff.Date{}
	err := date.SetDay(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1994)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("1st Day: %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}
