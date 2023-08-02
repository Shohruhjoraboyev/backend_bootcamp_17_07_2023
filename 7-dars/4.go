package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Printf("Son kiriting: ")
	fmt.Scan(&n)
	start := time.Now()
	for i := 0; i <= n; i++ {
		if tub(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	duration := time.Since(start)
	fmt.Println(duration)
}

func tub(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
