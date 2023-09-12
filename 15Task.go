/*
Exercise 15: Create a program that uses goroutines
and channels to implement a simple producer-consumer pattern.
The producer goroutine should generate numbers and send them to
the consumer goroutine through a channel, which will print the received numbers.
*/
package main

import "fmt"

func main() {
	numbers := make(chan int)
	done := make(chan bool)

	go producer(numbers)
	go consumer(numbers, done)
	<-done
}

func producer(numbers chan<- int) {
	for i := 1; i <= 10; i++ {
		numbers <- i
	}
	close(numbers)
}
func consumer(numbers chan int, done chan bool) {
	recived := []int{}
	for num := range numbers {
		recived = append(recived, num)
		// fmt.Println("Received: ", num)

	}
	fmt.Println("Received: ", recived)

	done <- true
}
