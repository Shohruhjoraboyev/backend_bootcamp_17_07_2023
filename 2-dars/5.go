package main

import "fmt"

func main() {
	var n int
	fmt.Print("n = ")
	fmt.Scan(&n)
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
