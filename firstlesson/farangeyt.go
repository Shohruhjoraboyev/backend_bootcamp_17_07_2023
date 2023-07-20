package main

import "fmt"

func main() {
	var t int
	fmt.Println("Farangeyt qiymatni kiriting: ")
	fmt.Scanln(&t)

	tk := t + 273
	tf := 5/9*t + 32
	t = (tk - 32) * 5 / 9
	fmt.Println(tk)
	fmt.Println(tf)
}
