package main

import "fmt"

func main() {
	arr := []int{5, 0, 1, 2, 3, 4}
	res := separateByIndex(arr)
	fmt.Println(res)
}

func separateByIndex(arr []int) []int {
	res := []int{}
	for _, v := range arr {
		res = append(res, arr[v])
	}
	return res
}
