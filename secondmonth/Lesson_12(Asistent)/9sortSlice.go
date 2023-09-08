/*
Exercise 9: Create a program that uses goroutines
and channels to sort a given slice of integers concurrently.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 4, 5, 3}
	ch := make(chan []int)
	go GetsortNumb(numbers, ch)
	sortedNumbers := <-ch
	fmt.Println(sortedNumbers)
}

func GetsortNumb(numbers []int, ch chan []int) {
	sort.Ints(numbers)

	ch <- numbers

	close(ch)
}
