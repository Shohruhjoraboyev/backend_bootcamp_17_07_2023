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

// Task #6
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	a := 1
	b := 1

	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b

}

// Task #7
func reverseNumbers(num int) int {
	res := 0

	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num = num / 10

	}
	return res

}

// Task #8
func birthdayCakeCandles(candles []int32) int32 {

	count := 0
	max := candles[0]

	for _, v := range candles {
		if v > max {
			max = v
		}
	}

	for _, v := range candles {
		if v == max {
			count++
		}
	}

	return int32(count)

}

// Task #9
func compareTriplets(a []int32, b []int32) []int32 {
	var x, y int32 = 0, 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			x += 0
		} else if a[i] > b[i] {
			x++
		} else {
			y++
		}

	}

	arr := []int32{x, y}

	return arr

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

	// // Task #5
	// fmt.Println(quadraticEquation(1, 4, -21)) // 1 -9

	// // Task #6
	// fmt.Println(fib(8)) // 1 -9

	// // Task #7
	// fmt.Println(reverseNumbers(81259))

	// // Task #8
	// candles := []int32{4, 2, 1, 4}
	// fmt.Println(birthdayCakeCandles(candles))

	// Task #9

	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))

}
