package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 2
	c := 5
	ckvadrat := c * c

	if ckvadrat == a*a+b*b {
		fmt.Println("Tog'ri burchakli uchburchak: ")
	} else {
		fmt.Println("Tog'ri burchakli uchburchak emas")
	}

}
