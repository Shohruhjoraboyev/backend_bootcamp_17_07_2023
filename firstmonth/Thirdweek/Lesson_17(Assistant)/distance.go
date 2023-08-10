package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(minDistance(8))
	// fmt.Println(minDistance(25))
	fmt.Println(minDistance(13013))
	// fmt.Println(minDistance(218683))
}

func minDistance(n int) int {
	factors := []int{}
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}

	minDistance := 0

	for i := 0; i < len(factors)-1; i++ {
		for j := i + 1; j < len(factors); j++ {
			distance := int(math.Abs(float64(factors[i] - factors[j])))
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}
