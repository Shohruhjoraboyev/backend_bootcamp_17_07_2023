/*
Exercise 10: Create a program that uses goroutines 
and channels to find the sum of squares of numbers from 1 to 10 concurrently.
*/

package main

import (
	"fmt"
)

func main() {
	n := 10
	ch := make(chan int)
	go SquaresNumb(n, ch)

	sum := 0
	for v := range ch {
		sum += v
	}

	fmt.Println("Sum of squares:", sum)
}

func SquaresNumb(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		kvadrat := i * i
		ch <- kvadrat
	}
	close(ch)
}
