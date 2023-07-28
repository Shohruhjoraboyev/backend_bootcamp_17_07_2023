package main

import "fmt"

func main() {
	text := "lenovo"
	var reverse string

	for _, r := range text {
		reverse = string(r) + reverse
	}
	defer fmt.Println("After reversed: ", reverse)
	fmt.Println("Before reversed: ", text)

}
