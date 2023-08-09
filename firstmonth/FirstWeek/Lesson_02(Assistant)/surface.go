package main

import "fmt"

func main() {
	var a int
	fmt.Println("Please enter number")
	fmt.Scanln(&a)

	fmt.Println("Kvadrat Surface : ", surface(a))

}
func surface(a int) int {
	s := a * a
	return s
}
