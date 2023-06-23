package main

import "fmt"

var p1 = fmt.Println

// defined types
type TeaSpoon float64
type TableSpoon float64
type ML float64

// defining conversions using previous defined types
func teaSpoonToML(tsp TeaSpoon) ML {
	return ML(tsp * 4.92)
}

func tableSpoonToML(tsp TableSpoon) ML {
	return ML(tsp * 14.79)
}

// defining associate methods
func (tsp TeaSpoon) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TableSpoon) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	// ML instance with TeaSpoon type
	ml1 := ML(TeaSpoon(3) * 4.92)
	fmt.Printf("3 teaspoons = %.2f mL\n", ml1)

	// ML instance with TableSpoon type
	ml2 := ML(TableSpoon(3) * 14.79)
	fmt.Printf("3 tablespoons = %.2f mL\n", ml2)

	// operating spoon types
	p1("2 teaspoons + 4 teaspoons =", TeaSpoon(2)+TeaSpoon(4))
	p1("2 teaspoons < 4 teaspoons =", TeaSpoon(2) < TeaSpoon(4))

	// using conversion methods
	fmt.Printf("3 teaspoons = %.2f mL\n", teaSpoonToML(3))
	fmt.Printf("3 teaspoons = %.2f mL\n", tableSpoonToML(3))

	tsp := TeaSpoon(1)
	tbs := TableSpoon(1)

	// using associate methods
	p1("1 teaspoon to ML =", tsp.ToMLs())
	p1("1 tablespoon to ML =", tbs.ToMLs())
}
