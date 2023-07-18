package main

import "fmt"

// Task #1
func oddOrEven(num int) string {
	if num%2 == 0 {
		return "Even"
	}

	return "Odd"

}

// Task #2
func swapNums(num1, num2 *int) {
	var num3 int
	num3 = *num1
	*num1 = *num2
	*num2 = num3

}

func main() {
	fmt.Println("Hello Go")

	fmt.Println(oddOrEven(4)) // Even
	fmt.Println(oddOrEven(5)) // Odd

	num1 := 3
	num2 := 5

	fmt.Println(num1)
	fmt.Println(num2)

	swapNums(&num1, &num2)

	fmt.Println(num1)
	fmt.Println(num2)

}
