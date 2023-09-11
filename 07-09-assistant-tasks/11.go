package main

// import (
// 	"fmt"
// )

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; x <= n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	var n int
// 	fmt.Print("n: ")
// 	fmt.Scan(&n)
// 	c := make(chan int)

// 	go fibonacci(n, c)

// 	for num := range c {
// 		fmt.Println(num)
// 	}
// }
