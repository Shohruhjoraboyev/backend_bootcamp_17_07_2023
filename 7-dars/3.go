package main

import "fmt"

func main() {
	adder := qosh()
	fmt.Println(adder(3))
	fmt.Println(adder(7))
	fmt.Println(adder(10))
}

func qosh() func(n int) int {
	son := 0
	return func(n int) int {
		son += n
		return son
	}
}
