package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	entries, err := os.ReadDir("example")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())

		// TODO: make this buffered - we only care about the imports which are at the beginning of the files
		body, err := os.ReadFile("example/" + e.Name())
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		fmt.Println("File: ", e.Name())
		fmt.Println(string(body))
		fmt.Println("")
	}
}
