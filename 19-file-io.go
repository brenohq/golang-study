package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var p1 = fmt.Println

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	// defer file closing
	defer f.Close()

	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string

	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// writing sPrimeArr to data.txt
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	// opening data.txt
	f, err = os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// reading data.txt content
	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		p1("Prime:", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
