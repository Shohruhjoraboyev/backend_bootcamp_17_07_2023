package main

import "fmt"

// Task #1
func oddOrEven(num int) string {
	if num%2 == 0 {
		return "Even"
	}

	return "Odd"

}

func main() {
	fmt.Println("Hello Go")

	fmt.Println(oddOrEven(4)) // Even
	fmt.Println(oddOrEven(5)) // Odd

}
