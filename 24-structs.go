package main

import "fmt"

var p1 = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAddress(c *customer, address string) {
	c.address = address
}

func main() {
	var tS customer

	tS.name = "Breno"
	tS.address = "England Street"
	tS.bal = 234.56

	getCustInfo(tS)
	newCustAddress(&tS, "South Street")

	p1("Address:", tS.address)

	sS := customer{name: "Ana Luiza", address: "Hungary Street", bal: 260.23}
	p1("Name:", sS.name)
	p1("Address:", sS.address)
	p1("Balance:", sS.bal)
}
