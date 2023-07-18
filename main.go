package main

import (
	"fmt"
	"sort"
)

func main() {
	// oddEven()
	// swapTwoNumbers()
	// sumOfMinAndMax()
	// sumOfMinAndMax2()

}

// odd or even
func oddEven() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println(num, "is even number")
	} else {
		fmt.Println(num, "is odd number")
	}
}

// swap two numbers
func swapTwoNumbers() {
	var a int
	var b int
	fmt.Print("Enter a value for a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter a value for b: ")
	fmt.Scanln(&b)

	fmt.Println("Before swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a

	fmt.Println("After swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func sumOfMinAndMax() {
	// var nums [3]int
	// fmt.Println("Enter a number for num 1: ")
	// fmt.Scanln(&nums[0])
	// fmt.Println("Enter a number for num 2: ")
	// fmt.Scanln(&nums[1])
	// fmt.Println("Enter a number for num 3: ")
	// fmt.Scanln(&nums[2])
	nums := [3]int{7, 1, 4}
	sort.Ints(nums[:])

	min := nums[0]
	max := nums[2]

	sum := min + max
	fmt.Println("sum of min and max of three numbers is =>", sum)
}

func sumOfMinAndMax2() {
	var num1, num2, num3, min, max int

	fmt.Println("Enter a number for num 1: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter a number for num 2: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter a number for num 3: ")
	fmt.Scanln(&num3)

	if num1 >= num2 && num1 >= num3 {
		max = num1
		if num2 >= num3 {
			min = num3
		} else {
			min = num2
		}
	} else if num2 >= num1 && num2 >= num3 {
		max = num2
		if num1 >= num3 {
			min = num3
		} else {
			min = num1
		}
	} else {
		max = num3
		if num1 >= num2 {
			min = num2
		} else {
			min = num1
		}
	}

	sum := min + max

	fmt.Println("sum of min and max of three numbers is =>", sum)

}
