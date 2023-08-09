package main

import "fmt"

func main() {
	number := []int{234567}
	reversed := []int{}

	for i := len(number); i >= 0; i-- {
		reversed = append(reversed, i)

	}
	fmt.Println(reversed)
}
