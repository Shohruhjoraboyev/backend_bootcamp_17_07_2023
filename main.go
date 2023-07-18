package main

import (
	"fmt"
	"math"
)

// Task #1
func oddOrEven(num int) string {
	if num%2 == 0 {
		return "Even"
	}

	return "Odd"

}

// Task #2
func swapNums(num1, num2 *int) {
	var num3 int
	num3 = *num1
	*num1 = *num2
	*num2 = num3

}

// Task #3
func sumMaxMin(a, b, c float64) float64 {
	var max float64
	var min float64

	//find max
	max = a

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	// find min
	min = a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return max + min

}

// Task #4
func distanceOfTwoPoints(x1, y1, x2, y2 float64) float64 {
	x := math.Pow(x2-x1, 2)
	y := math.Pow(y2-y1, 2)
	return math.Sqrt(x + y)

}

// Task #5
func quadraticEquation(a, b, c float64) (float64, float64) {
	x1 := -float64(b) + math.Sqrt(math.Pow(b, 2)-4*a*c)/2*a
	x2 := -float64(b) - math.Sqrt(math.Pow(b, 2)-4*a*c)/2*a

	return x1, x2

}

func main() {
	fmt.Println("Hello Go")

	// // Task #1
	// fmt.Println(oddOrEven(4)) // Even
	// fmt.Println(oddOrEven(5)) // Odd

	// // Task #2
	// num1 := 3
	// num2 := 5

	// fmt.Println(num1)
	// fmt.Println(num2)

	// swapNums(&num1, &num2)

	// fmt.Println(num1)
	// fmt.Println(num2)

	// // Task #3
	// fmt.Println(sumMaxMin(1, 3, 2)) // 4
	// fmt.Println(sumMaxMin(4, 5, 6)) // 10

	// // Task #4
	// fmt.Println(distanceOfTwoPoints(5, 4, 1, 1)) // 5

	// Task #5
	fmt.Println(quadraticEquation(1, 4, -21)) // 1 -9
}
