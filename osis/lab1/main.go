package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func writeFile(files []string) {
	file, err := os.OpenFile("filenames.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		file.WriteString(f + "\n")
	}
}

func main() {
	var files []string

	root := "/home/myxtar/Downloads"

	fmt.Println("Enter extension: ")

	var extGlob string
	fmt.Scan(&extGlob)

	fmt.Println("Scanning..." + "\n")

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(path)
			if ext == extGlob {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(files)
	writeFile(files)
}
