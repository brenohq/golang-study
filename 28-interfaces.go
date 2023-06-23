package main

import "fmt"

var p1 = fmt.Println

// interface declaration
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

// binding method to Cat type
func (c Cat) Attack() {
	p1("Cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

// implementing interface methods for Cat type
func (c Cat) AngrySound() {
	p1("Cat says Hissss")
}

func (c Cat) HappySound() {
	p1("Cat says Purrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	kitty.HappySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	p1("Cats name:", kitty2.Name())
}
