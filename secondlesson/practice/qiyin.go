package main

import "fmt"

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	res := []int{}
	arr1 := arr[:3]
	arr2 := arr[3:]

	for i := 0; i < 3; i++ {

		res = append(res, arr1[i], arr2[i])
		// res = append(res, arr2[i])
	}
	fmt.Println(res)

}
