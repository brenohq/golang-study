package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// removing filename default arg on first position
	args := os.Args[1:]

	var iArgs = []int{}

	// iterating through the args array and appending parsed values to iArgs variable
	for _, i := range args {
		val, err := strconv.Atoi(i)

		if err != nil {
			panic(err)
		}

		iArgs = append(iArgs, val)
	}

	max := 0

	// finding max value present in iArgs
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}

	fmt.Println("Max value:", max)
}
