package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Please enter kvadrat first side: ")
	fmt.Scanln(&a)
	fmt.Println("Please enter kvadrat second side: ")
	fmt.Scanln(&b)

	area, perimeter := surfaceKvadrat(a, b)
	fmt.Println("Kvadrat area: ", area)
	fmt.Println("Kvadrat perimeter: ", perimeter)
}

func surfaceKvadrat(a, b int) (int, int) {
	p := 2 * (a + b)
	s := a * b
	return s, p
}
