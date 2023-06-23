package main

import "fmt"

var p1 = fmt.Println

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.height * r.length
}

func main() {
	rect := rectangle{height: 25.2, length: 12.2}

	p1("Rect area:", rect.Area())
}
