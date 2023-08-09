package main

import "fmt"

func main() {
	slice := []int{2, 3, 4, 5, 7, 8, 9, 11, 23}
	low := 4
	high := 12

	result := make([]int, 0)
	for _, num := range slice {
		if num >= low && num <= high {
			result = append(result, num)
		}
	}

	fmt.Println(result)
}
