package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	mySl := []int{1, 5, 3, 7, 12, 4, 2}
// 	ch := make(chan int)
// 	go task9(mySl, ch)
// 	for range mySl {
// 		fmt.Printf("%d ", <-ch)
// 	}

// 	fmt.Println()
// }

// func task9(sl []int, ch chan int) {
// 	sort.Ints(sl)
// 	for _, v := range sl {
// 		ch <- v
// 	}
// }
