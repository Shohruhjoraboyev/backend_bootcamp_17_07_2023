/*
Exercise 1: Create a program that uses goroutines 
and channels to print numbers from 1 to 10 concurrently.
*/
package main

import (
	"fmt"
)

func main() {
	n := 10
	ch := make(chan int)
	go OneToTen(n, ch)


	for v := range (ch) {
		if v <= n {
			fmt.Println(v)
		}

	}

}

func OneToTen(n int, ch chan int) {

	for i := 0; i < n; i++ {
		ch <- i

	}
	close(ch)
}
