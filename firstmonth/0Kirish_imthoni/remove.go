package main

import "fmt"

func main() {
	input := "sa81l3o2m3"
	fmt.Println(getNum(input))
}

func getNum(str string) []string {
	result := []string{}
	numbers := ""

	for _, v := range str {
		if v >= '0' && v <= '9' {
			numbers += string(v)
		}
	}

	result = append(result, numbers)

	return result
}
