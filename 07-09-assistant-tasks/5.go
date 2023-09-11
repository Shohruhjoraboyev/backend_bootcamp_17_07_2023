package main

// import "fmt"

// func main() {
// 	mySl := []int{1, 5, 3, 7, 12, 4, 2}
// 	ch := make(chan int, 1)
// 	go task5(mySl, ch)
// 	fmt.Println(<-ch)
// }

// func task5(sl []int, ch chan int) {
// 	max := sl[0]
// 	for _, v := range sl {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	ch <- max
// }
