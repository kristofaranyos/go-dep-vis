package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	depMap, err := getDependencyMap("example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(depMap)
}

func getDependencyMap(folderName string) (DependencyMap, error) {
	entries, err := os.ReadDir(folderName)
	if err != nil {
		return nil, fmt.Errorf("could not read folder %q: %w", folderName, err)
	}

	result := make(DependencyMap, 0)

	for _, e := range entries {
		path := fmt.Sprintf("%s/%s", folderName, e.Name())

		if e.IsDir() {
			m, err := getDependencyMap(path)
			if err != nil {
				return nil, err
			}

			if len(m) > 0 {
				result.Merge(m)
			}

			continue
		}

		if !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
		if err != nil {
			panic(err)
		}

		var list []string
		for _, spec := range file.Imports {
			list = append(list, spec.Path.Value)
		}

		if len(list) > 0 {
			result.Add(path, list)
		}
	}

	return result, nil
}
