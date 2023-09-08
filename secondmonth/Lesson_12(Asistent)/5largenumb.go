/*
Exercise 5: Create a program that uses goroutines 
and channels to find the largest number in a given slice concurrently.
*/
package main

import "fmt"

func main() {
	numbers := []int{23, 45, 52, 34}
	ch := make(chan int)
	go getLargeNumb(numbers, ch)
	fmt.Println(<-ch)
}

func getLargeNumb(numbers []int, ch chan int) {
	max := 0 

	for _, v := range numbers {

		if max <= v {
			max = v
		}
	}

	ch <- max
	close(ch)
}
