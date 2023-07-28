package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 1}
	var count int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++

			}
		}
	}
	fmt.Println(count)
}
