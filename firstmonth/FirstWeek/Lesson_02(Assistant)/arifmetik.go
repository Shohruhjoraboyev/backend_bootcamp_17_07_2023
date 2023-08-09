package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := 3, 2, 5

	arifmetik, geometrik := arifmetik(a, b, c)
	fmt.Println("O'rta arifmetik: ", arifmetik)
	fmt.Println("O'rta geometrik: ", geometrik)
}

func arifmetik(a, b, c int) (float64, float64) {
	d := (a + b + c) / 3

	g := math.Pow(float64(a*b*c), 1./3)

	return float64(d), g

}
