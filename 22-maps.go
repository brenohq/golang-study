package main

import "fmt"

var p1 = fmt.Println

func main() {
	// creating variable
	var heroes map[string]string

	// allocating with = and with :=
	heroes = make(map[string]string)
	villains := make(map[string]string)

	// defining key/values
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villains["Lex Luthor"] = "Lex Luthor"

	// defining with literals
	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	// acessing key/value
	fmt.Printf("Batman identity is %v\n", heroes["Batman"])

	// detecting nil values
	_, ok := superPets[3]
	p1("Chip: ", superPets[3])
	p1("Is there a 3rd pet:", ok)

	// iterating throught the map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// deleting one key/value
	delete(heroes, "The Flash")
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
}
