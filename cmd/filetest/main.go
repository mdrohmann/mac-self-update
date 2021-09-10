package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.log", os.O_TRUNC|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not open file ./test.log: %v", err)
	}

	if err := os.Remove("./test.log"); err != nil {
		log.Fatalf("Failed to remove ./test.log: %v", err)
	}

	if _, err := f.WriteString("some text"); err != nil {
		log.Fatalf("Failed to write some text: %v", err)
	}
}
