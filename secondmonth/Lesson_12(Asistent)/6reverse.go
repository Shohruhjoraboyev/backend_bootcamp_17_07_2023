/*
Exercise 6: Create a program that uses goroutines 
and channels to reverse a string concurrently.
*/

package main

import "fmt"

func main() {
	sample := "Password"
	ch := make(chan string)
	go reversed(sample, ch)
	fmt.Println(<-ch)
}

func reversed(sample string, ch chan string) {
	reversedStr := ""
	for i := len(sample) - 1; i >= 0; i-- {
		reversedStr += string(sample[i])
	}
	ch <- reversedStr
	close(ch)
}