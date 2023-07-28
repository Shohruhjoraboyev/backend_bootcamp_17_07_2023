package main

import (
	"fmt"
	ouf "github/jdu90009/firtslesson/pifagor"
)

func main() {
	ouf.Pifagor()
	a, b, c := 3, 15, 23
	minimum := a
	maximum := a

	if b < minimum {
		minimum = b
	}

	if c < minimum {
		minimum = c
	}

	if b > maximum {
		maximum = b
	}

	if c > maximum {
		maximum = c

	}

	fmt.Println("Sum of 3 numbers is:", maximum+minimum)

}
