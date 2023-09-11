package main

// import "fmt"

// func main() {
// 	mySl := []int{1, 5, 3, 7, 12, 4, 2}
// 	ch := make(chan float64)
// 	go task8(mySl, ch)
// 	fmt.Printf("%f\n", <-ch)

// }

// func task8(sl []int, ch chan float64) {
// 	sum := 0
// 	for _, v := range sl {
// 		sum += v
// 	}

// 	ch <- float64(sum) / float64(len(sl))
// }
