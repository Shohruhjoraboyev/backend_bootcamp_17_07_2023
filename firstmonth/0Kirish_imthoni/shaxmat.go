package main

import "fmt"

func check(x string, y int) {

	if (x == "a" || x == "c" || x == "e" || x == "g") && y%2 == 1 {
		fmt.Println("BLACK")
	} else if (x == "b" || x == "d" || x == "f" || x == "h") && y%2 == 0 {
		fmt.Println("BLACK")
	} else {
		fmt.Println("WHITE")
	}
}

func main() {

	check("a", 1)
	check("d", 5)
	check("d", 2)
}
