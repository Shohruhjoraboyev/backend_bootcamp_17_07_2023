package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 4, 3}
	max := arr[0]
	count := 0
	for _, v := range arr {
		if max < v {
			max = v
			count = 1
		} else if max == v {
			count++
		}
	}
	fmt.Println("Max numb: ", max)
	fmt.Println("count of max numb: ", count)

}
