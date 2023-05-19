package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var p1 = fmt.Println

func main() {
	// simple for loop
	for x := 1; x <= 5; x++ {
		p1(x)
	}

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	randNum := rand.Intn(50) + 1

	reader := bufio.NewReader(os.Stdin)

	// while true loop
	for true {
		fmt.Print("Guess a number between 0 and 50:\n")
		p1("Random number is:", randNum)
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			p1("Pick a lower value")
		} else if iGuess < randNum {
			p1("Pick a higher value")
		} else {
			p1("You guessed it!")
			break
		}

		p1("\n")
	}
}
