/*
3. Closure yordamida, men argumentga nechchi bersam shuni qo'shib qaytaradigan function yozing.
masalan:

	adder(3) => 3
	adder(7) =>10
*/
package main

import "fmt"

func main() {
	a := adder()
	fmt.Println(a(3))
	fmt.Println(a(7))
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
