package main

// import (
// 	"fmt"
// )

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func checkPrime(num int, ch chan<- bool) {
// 	ch <- isPrime(num)
// }

// func main() {
// 	num := 17

// 	ch := make(chan bool)
// 	go checkPrime(num, ch)
// 	res := <-ch
// 	if res {
// 		fmt.Printf("%d is prime\n", num)
// 	} else {
// 		fmt.Printf("%d is not prime\n", num)
// 	}
// }
