package main

import "fmt"

func main() {
	var a int
	fmt.Println("Please enter number")
	fmt.Scanln(&a)

	fmt.Println("Kvadrat Perimetri : ", perimetr(a))

}
func perimetr(a int) int {
	p := 4 * a
	return p
}
