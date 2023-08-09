package main

import "fmt"

func main() {
	arr := []string{"++x", "--x", "++x"}
	var sum int

	for _, v := range arr {
		if v == "++x" {
			sum++
		} else if v == "--x" {
			sum--
		} else {
			fmt.Println("Some bug has")
		}
	}
	fmt.Println(sum)

}
