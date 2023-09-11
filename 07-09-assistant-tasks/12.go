package main

// import (
// 	"fmt"
// )

// func main() {
// 	mySl := []int{1, 5, 3, 7, 12, 4, 2}
// 	ch := make(chan int)
// 	go task12(mySl, ch)

// 	fmt.Printf("%d\n", <-ch)
// }

// func task12(sl []int, ch chan int) {
// 	sum := 0
// 	for _, v := range sl {
// 		if v%2 == 0 {
// 			sum += v
// 		}
// 	}
// 	ch <- sum
// }
