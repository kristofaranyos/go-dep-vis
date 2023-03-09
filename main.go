package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(traverseFiles("example"))
}

func traverseFiles(folderName string) error {
	entries, err := os.ReadDir(folderName)
	if err != nil {
		return fmt.Errorf("could not read folder %q: %w", folderName, err)
	}

	for _, e := range entries {
		path := fmt.Sprintf("%s/%s", folderName, e.Name())

		if e.IsDir() {
			if err := traverseFiles(path); err != nil {
				return err
			}

			continue
		}

		// TODO: make this buffered - we only care about the imports which are at the beginning of the files
		body, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read file %q: %w", path, err)
		}

		fmt.Println("File: ", path)
		fmt.Println(string(body))
		fmt.Println("")
	}

	return nil
}
