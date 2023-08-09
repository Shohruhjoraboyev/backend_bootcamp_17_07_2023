package main

import (
	"fmt"
)

func main() {

	n := 4
	stair(n)
}

func stair(n int) {

	for i := 0; i < n; i++ {
		pattern := ""

		for j := 0; j <= n-i; j++ {
			pattern += " "

		}

		for j := 1; j <= i+1; j++ {
			pattern += "#"

		}
		fmt.Println(pattern)
	}

	1,2,3

	1,3
	

}
