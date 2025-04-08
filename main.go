package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	rootDirectory string
	fileName      string
)

func main() {
	flag.StringVar(&rootDirectory, "rootDirectory", "", "Root directory to find")
	flag.StringVar(&fileName, "fileName", "", "File name to find")
	flag.Parse()

	if len(fileName) < 1 {
		log.Fatal("No --fileName is given")
	}

	if len(rootDirectory) < 1 {
		log.Fatal("No --rootDirectory is given")
	}

	foundPath, err := FilePathWalkDir(rootDirectory, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(foundPath)
}

func FilePathWalkDir(root string, lookFor string) (string, error) {
	var foundPath string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if lookFor == info.Name() {
				foundPath = path
			}
		}
		return nil
	})
	return foundPath, err
}
