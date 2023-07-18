package main

import (
	"fmt"
	"math"
)

func main() {

	// odd or even

	var num = 0

	if num == 0 {
		fmt.Println(num, " is zero")
	} else if num%2 == 1 {
		fmt.Println(num, " is odd")
	} else {
		fmt.Println(num, " is even")
	}

	// swap to numbers

	var first = 5
	var second = 6

	first = first + second
	second = first - second
	first = first - second

	fmt.Println(" First Number = ", first, "\n", "Second Number = ", second)

	// summ of maximum and minimum of given 3 numbers

	var firsNumber, secondNumber, thirdNumber = 5, 2, 1

	var max, min int

	max = firsNumber
	min = secondNumber

	if max < secondNumber {
		min = firsNumber
		max = secondNumber
	}
	if max < thirdNumber {
		max = thirdNumber
	}
	if min > thirdNumber {
		min = thirdNumber
	}

	fmt.Println("max + min = ", max+min)

	//distance between two points({x1;y1} and {x2;y2}

	var x1, y1, x2, y2 = 1, -6, -2, -2
	var a, b int
	var c float64

	if (x1 < 0 && x2 < 0) || (x1 > 0 && x2 > 0) {

		a = int(math.Abs(math.Abs(float64(x1)) - math.Abs(float64(x2))))

	} else {

		a = int(math.Abs(float64(x1)) + math.Abs(float64(x2)))

	}

	if (y1 < 0 && y2 < 0) || (y1 > 0 && y2 > 0) {

		b = int(math.Abs(math.Abs(float64(y1)) - math.Abs(float64(y2))))

	} else {

		b = int(math.Abs(float64(y1)) + math.Abs(float64(y2)))

	}

	c = (math.Sqrt(float64(a*a + b*b)))

	fmt.Println("distance between two points({x1;y1} and {x2;y2} = ", c)

	//solution of quadratic equation

	var num1, num2, num3 = 2, 3, 4

	equation(num1, num2, num3)

}

func equation(a, b, c int) (float64, float64) {

	var x1 float64 = float64(-(float64(b) + math.Sqrt(float64(b*b)-float64(4*a*c))) / float64(2*a))
	var x2 float64 = float64((float64(b) + math.Sqrt(float64(b*b)-float64(4*a*c))) / float64(2*a))

	return x1, x2

}
