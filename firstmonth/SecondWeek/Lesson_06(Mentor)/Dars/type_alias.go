package structmap

import "fmt"

type integer int32

type arrayInt []int

func alias() {
	var arr arrayInt
	var b integer
	arr2 := arrayInt{5, 6, 7, 8, 9}
	b = 12
	arr = append(arr, 1, 2, 3)

	fmt.Println(b)
	fmt.Println(arr)
	fmt.Println(arr2)
}
