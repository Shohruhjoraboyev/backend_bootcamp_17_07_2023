/*
Exercise 13: Create a program that uses goroutines
and channels to find the length of the longest word in a given string concurrently.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Create a program that uses goroutines and channels to find the length of the longest word in a given string concurrently."
	ch := make(chan int)

	go GetLongestWordLength(str, ch)

	length := <-ch
	fmt.Println(length)
}

func GetLongestWordLength(str string, ch chan int) {
	words := strings.Fields(str)
	longestLength := 0

	for _, word := range words {
		if len(word) > longestLength {
			longestLength = len(word)
		}
	}

	ch <- longestLength
	close(ch)
}
