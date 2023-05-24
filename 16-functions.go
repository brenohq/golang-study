package main

import (
	"fmt"
)

var p1 = fmt.Println

// func sayHello() {
// 	p1("Hello!")
// }

// func getSum(x int, y int) int {
// 	return x + y
// }

// func getTwo(x int) (int, int) {
// 	return x, x + 2
// }

// func getQuocient(x float64, y float64) (ans float64, err error) {
// 	if y == 0 {
// 		return 0, fmt.Errorf("Division by zero")
// 	} else {
// 		return x / y, nil
// 	}
// }

// func getSumV2(nums ...int) int {
// 	sum := 0

// 	for _, num := range nums {
// 		sum += num
// 	}

// 	return sum
// }

// func getArraySum(arr []int) int {
// 	sum := 0

// 	for _, val := range arr {
// 		sum += val
// 	}

// 	return sum
// }

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeValV2(myPtr *int) {
	*myPtr = 12
}

func main() {
	// declaration example
	// func funcName(parameters) returnType { BODY }

	// simple function
	// sayHello()

	// 2 arguments
	// p1(getSum(5, 23))

	// 2 returns
	// p1(getTwo(5))

	// error throwing
	// p1(getQuocient(5, 0))
	// p1(getQuocient(5, 2))

	// dynamic arguments
	// p1(getSumV2(2, 3, 4, 5))

	// array argument
	// p1(getArraySum([]int{2, 3, 4, 5}))

	// passing argument by values
	f3 := 5
	p1("f3 before func:", f3)
	changeVal(f3)
	p1("f3 after func:", f3)

	// passing argument by reference
	f4 := 5
	p1("f4 before func:", f4)
	changeValV2(&f4)
	p1("f4 after func:", f4)
}
