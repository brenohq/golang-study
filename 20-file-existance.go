package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var p1 = fmt.Println

func main() {
	_, err := os.Stat("data.txt")

	if errors.Is(err, os.ErrNotExist) {
		p1("File doesn't exist")
	} else {
		// opening file in write mode
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		// writing new content to file
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
