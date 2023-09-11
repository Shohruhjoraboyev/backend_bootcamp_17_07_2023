package main

// import "fmt"

// func main() {
// 	ch := make(chan string)
// 	go task6("salom", ch)
// 	fmt.Println(<-ch)
// }

// func task6(s string, ch chan string) {
// 	ch <- reverse(s)
// }

// func reverse(s string) string {
// 	rns := []rune(s)
// 	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
// 		rns[i], rns[j] = rns[j], rns[i]
// 	}

// 	return string(rns)
// }
