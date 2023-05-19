package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var p1 = fmt.Println

func main() {
	// float precision
	p1("Float precision =", 0.111111111111111111111111+0.111111111111111111111111)

	// math operations
	p1("Abs(-10) =", math.Abs(-10))
	p1("Pow(4, 2) =", math.Pow(4, 2))
	p1("Sqrt(16) =", math.Sqrt(16))
	p1("Cbrt(8) =", math.Cbrt(8))
	p1("Ceil(4.4) =", math.Ceil(4.4))
	p1("Floor(4.4) =", math.Floor(4.4))
	p1("Round(4.4) =", math.Round(4.4))
	p1("Log2(8) =", math.Log2(8))
	p1("Log10(100) =", math.Log10(100))
	p1("Log(7.389) =", math.Log(math.Exp(2)))
	p1("Max(5, 4)", math.Max(5, 4))
	p1("Min(5, 4)", math.Min(5, 4))

	// random number
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	randNum := rand.Intn(50) + 1
	p1("Random number:", randNum)

	// radians degrees
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)

	// sin/cos
	p1("Sin(90) =", math.Sin(r90))
	p1("Cos(90) =", math.Cos(r90))
}
