package main

import (
	"fmt"
	"time"
)

var p1 = fmt.Println

func main() {
	now := time.Now()
	p1(now.Year(), now.Month(), now.Day())
	p1(now.Hour(), now.Minute(), now.Second())
}
