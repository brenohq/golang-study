package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var p1 = fmt.Println

func main() {
	p1("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		p1("Hello", name)
	} else {
		log.Fatal(err)
	}
}
