package main

import "fmt"

func main() {
	var s = "ubuntu"

	defer fmt.Println()
	for _, v := range s {
		defer fmt.Print(string(v))
	}
}
