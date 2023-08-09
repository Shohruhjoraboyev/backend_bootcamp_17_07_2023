package main

import "fmt"

func main() {
	a := []int{34, 23, 15, 12}
	b := []int{34, 43, 15, 12}
	bob := 0
	alice := 0
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] < v {
				alice++

			} else if b[i] > v {
				bob++
			}
		}
	}
	fmt.Println("Alice's score: ", alice)
	fmt.Println("Bob's score: ", bob)
}

//  func compare( a[] int , b[] int)(int,int){
// 	if len(a) == len(b) {
// 		for i, v := range a {
// 			if b[i] < v {
// 				alice++

// 			} else if b[i] > v {
// 				bob++
// 			}
// 		}
// 	}
// 	fmt.Println("Alice's score: ", alice)
// 	fmt.Println("Bob's score: ", bob)

//  }
