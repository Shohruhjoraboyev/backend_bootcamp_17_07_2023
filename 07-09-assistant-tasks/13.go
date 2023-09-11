package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	ch := make(chan int)
// 	s := "mening ismim Diyorbek Hasanov"
// 	go task13(s, ch)

// 	fmt.Printf("%d\n", <-ch)
// }

// func task13(s string, ch chan int) {
// 	mySl := strings.Split(s, " ")

// 	maxString := mySl[0]
// 	for _, v := range mySl {
// 		if len(v) > len(maxString) {
// 			maxString = v
// 		}
// 	}

// 	ch <- len(maxString)
// }
