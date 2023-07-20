package main

import "fmt"

func main() {
	arr := []int{5, 0, 1, 2, 3, 4}
	res := []int{}

	for _, v := range arr {
		res = append(res, arr[v])

	}
	fmt.Println(res)

}
