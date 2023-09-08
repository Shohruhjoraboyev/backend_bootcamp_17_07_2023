package main

import (
	"fmt"
)

/*
Exercise 8: Create a program that uses goroutines 
and channels to find the average of numbers in a given slice concurrently.
*/
func main() {

	numbers := []int{2, 4, 5, 3}
	ch := make(chan float32)
	go getAvg(numbers, ch)
	fmt.Println(<-ch)
}

func getAvg(numbers []int, ch chan float32) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	ch <- float32(sum / len(numbers))
	close(ch)
}
