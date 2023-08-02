package main

import "fmt"

func main() {
	fibo := fibonachi()

	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}

}

func fibonachi() func() int {
	var x int = 0
	var y int = 1
	return func() int {
		ret := x
		x, y = y, x+y
		return ret
	}
}
