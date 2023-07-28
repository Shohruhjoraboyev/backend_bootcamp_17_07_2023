package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Println("Please enter kub side: ")
	fmt.Scanln(&a)

	hajm, sirti := surfaceKub(a)
	fmt.Println("Kub hajmi: ", hajm)
	fmt.Println("Kub sirti: ", sirti)
}

func surfaceKub(a int) (float64, float64) {
	v := math.Pow(float64(a), 3)
	s := 6 * math.Pow(float64(a), 2)
	return v, s

}
