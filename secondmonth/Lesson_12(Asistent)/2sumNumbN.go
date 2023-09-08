/*
Exercise 2: Create a program that uses goroutines 
and channels to calculate the sum of numbers from 1 to 100 concurrently.
*/

package main

import (
	"fmt"
)

func main() {
	n := 100
	ch := make(chan int)
	go SumN(n, ch)

	sum := <-ch 

	fmt.Println("Sum of numbers from 1 to", n, "is", sum)
}

func SumN(n int, ch chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	ch <- sum 
	close(ch)
}
