package main

import "fmt"

func main() {
	num1 := 3
	num2 := -78
	var a *int
	var b *int
	a = &num1
	b = &num2
	*a, *b = *b, *a

	fmt.Println("Swapped values:")
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)
}
