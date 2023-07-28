package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{{1, 2, 3}, {2, 5, 6}}
	max:=math.Max(...arr)
	fmt.Println(max)

	arr1 := arr[:1]
	arr2 := arr[1:]
	sum1 := 0
	sum2 := 0

	for _, v := range arr1[0] {
		sum1 += v
	}
	for _, v := range arr2[0] {
		sum2 += v
	}
	if sum1 > sum2 {
		fmt.Println(sum1)
	} else {
		fmt.Println(sum2)
	}

}
