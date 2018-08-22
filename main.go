package main

import (
	"fmt"
)

func main() {

	comparemapresult := comparemap()

	if len(comparemapresult) > 0 {
		for k1, k2 := range comparemapresult {
			fmt.Printf("%s file exist here : %s\n", k1, k2)
		}
	} else {
		fmt.Println("No files are the same in both paths")
	}
}
