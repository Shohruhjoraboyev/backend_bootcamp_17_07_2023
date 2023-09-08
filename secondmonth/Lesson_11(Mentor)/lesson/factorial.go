package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 5
	ch := make(chan int)
	go factorial(n, ch)
	result := <-ch
	fmt.Println("Ushbu sonning Factoriali:", n, "--", result)
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
	close(ch)
}
