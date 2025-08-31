package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	rootDirectory string
	fileName      string
)

func main() {
	flag.StringVar(&rootDirectory, "rootDirectory", "", "Root directory to find")
	flag.StringVar(&fileName, "fileName", "", "File name to find, regex pattern eg. \".*\.cmake\" supported")
	flag.Parse()
	
	if len(fileName) < 1 {
		log.Fatal("No --fileName is given")
	}

	if len(rootDirectory) < 1 {
		log.Fatal("No --rootDirectory is given")
	}

	matchedPathArray, err := FilePathWalkDir(rootDirectory, fileName)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, matchedPath := range matchedPathArray {
    	fmt.Println(matchedPath)
    }
}

func stringMatch(patternTolookFor string, content string) bool {
	match, _ := regexp.MatchString(patternTolookFor, content)
	return match
}

func FilePathWalkDir(root string, lookFor string) ([]string, error) {
	var matchedPathArray []string
	
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			name := info.Name()
			if lookFor == name || stringMatch(lookFor, name){
				matchedPathArray = append(matchedPathArray, path)
				// DO NOT RETURN
			}
		}
		return nil
	})
	return matchedPathArray, err
}
