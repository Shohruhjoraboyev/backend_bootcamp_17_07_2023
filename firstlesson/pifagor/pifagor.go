package pifagor

import (
	"fmt"
	"math"
)

func Pifagor() {
	a, b, c := 1.0, -3.0, 2.0
	discriminant := b*b - 4*a*c

	radix := math.Sqrt(discriminant)
	X1 := (-b + radix) / (2 * a)
	X2 := (-b - radix) / (2 * a)
	fmt.Printf("The solutions are x1 = %.2f and x2 = %.2f\n", X1, X2)

}
