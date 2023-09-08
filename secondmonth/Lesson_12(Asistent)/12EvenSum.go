/*
Exercise 12: Create a program that uses goroutines 
and channels to find the sum of even numbers in a given slice concurrently.
*/

package main

import "fmt"

func main() {
	evenSlice := []int{12, 21, 32, 43, 34}
	ch := make(chan int)

	go GetSumEvenNumb(evenSlice, ch)
	fmt.Println(<-ch)

}

func GetSumEvenNumb(evenSlice []int, ch chan int) {
	sum := 0

	for _, v := range evenSlice {
		if v%2 == 0 {
			sum = sum + v
		}
	}
	ch <- sum
	close(ch)
}
