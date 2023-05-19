package main

import (
	"fmt"
)

var p1 = fmt.Println

func main() {
	var arr1 [5]int
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}

	p1("Arr1 Index 0:", arr1[0])
	p1("Arr2 length:", len(arr2))

	// default for
	for i := 0; i < len(arr2); i++ {
		p1(arr2[i])
	}

	// range for
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// multi dimensional arrays
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// matrix for
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			p1(arr3[i][j])
		}
	}

	// string to rune array
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune array: %d\n", v)
	}

	// byte array
	byteArr := []byte{'a', 'b', 'c', 'd'}
	bStr := string(byteArr[:])
	p1("I'm a string:", bStr)
}
