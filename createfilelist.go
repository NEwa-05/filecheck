package main

import (
	"os"
	"path/filepath"
)

func createfilelist() []string {
	fileList := make([]string, 0)
	for _, element := range getpathfromuser() {
		e := filepath.Walk(element, func(path string, f os.FileInfo, err error) error {
			unusedir, err := os.Stat(path)
			if unusedir.IsDir() {
			} else {
				fileList = append(fileList, path)
			}
			return err
		})
		if e != nil {
			panic(e)
		}
	}
	return fileList
}
