package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		num := rand.Intn(500)
		ch <- num
		time.Sleep(500 * time.Millisecond)
	}

	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
