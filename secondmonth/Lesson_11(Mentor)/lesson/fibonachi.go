package main

import (
	"fmt"
)

func main() {
	n := 10
	ch := make(chan int)
	go fibonachi(n, ch)

	for v := range ch {
		if v <= n {
			fmt.Println(v)
		}

	}

}

func fibonachi(n int, ch chan int) {
	a, b := 1, 1

	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}
