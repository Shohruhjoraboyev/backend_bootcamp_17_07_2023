/*
Exercise 4: Create a program that uses goroutines 
and channels to check if a given number is prime concurrently.
*/
package main

import (
	"fmt"
)

func main() {
	n := 7
	ch := make(chan int)
	go primeCheck(n, ch)
	result := <-ch

	if n == result {
		fmt.Println(result, "This is a prime number")
	} else {
		fmt.Println("This is not a prime number")
	}
}

func primeCheck(n int, ch chan int) {
	if n <= 1 {
		ch <- n
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				ch <- n
				break
			}
		}
	}

	close(ch)
}
