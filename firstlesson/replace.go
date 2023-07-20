package main

import "fmt"

func main() {
	num1 := 5
	num2 := 10

	a := num1
	num1 = num2
	num2 = a

	fmt.Println("Swapped values:")
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

}
