package main

import "fmt"

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	arr1 := arr[:3]
	arr2 := arr[3:]

	fmt.Println(seperateByIndex(arr1, arr2))

}
func seperateByIndex(arr1, arr2 []int) []int {
	res := []int{}

	for i := 0; i < 3; i++ {
		res = append(res, arr1[i], arr2[i])
		// res = append(res, arr2[i])
	}
	return res
}
