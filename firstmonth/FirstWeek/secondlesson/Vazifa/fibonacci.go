package main

import "fmt"

func main() {
	son := 5
	fmt.Println("Fibonachi: ", fibonaci(son))

}
func fibonaci(n int) int {
	var x, y, z = 1, 1, 0
	for i := 0; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}
	return z
}
