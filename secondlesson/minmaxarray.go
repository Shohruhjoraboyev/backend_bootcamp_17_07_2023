package main

import "fmt"

func main() {
	arr := []int{34, 23, 15, 12}
	min := arr[0]
	max := arr[0]
	sum := 0
	for _, v := range arr {
		sum += v
		if max < v {
			max = v
		} else if min > v {
			min = v
		}

	}
	fmt.Println("Max numb: ", max)
	fmt.Println("Min numb: ", min)
	fmt.Println("Sum numb: ", sum/len(arr))
}
