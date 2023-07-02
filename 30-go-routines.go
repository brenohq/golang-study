package main

import (
	"fmt"
)

var p1 = fmt.Println

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	p1(<-channel1)
	p1(<-channel1)
	p1(<-channel1)

	p1(<-channel2)
	p1(<-channel2)
	p1(<-channel2)
}
