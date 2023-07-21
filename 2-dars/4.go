package main

import "fmt"

func main() {
	var n, res int
	fmt.Print("n = ")
	fmt.Scan(&n)

	for n > 0 {
		res = res*10 + n%10
		n /= 10
	}

	fmt.Println(res)
}
