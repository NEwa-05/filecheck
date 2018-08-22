package main

import (
	"fmt"
)

func getpathfromuser() []string {
	userPath := make([]string, 0)
	fmt.Print("Enter path: ")
	var first string
	fmt.Scanln(&first)
	userPath = append(userPath, first)
	return userPath
}
