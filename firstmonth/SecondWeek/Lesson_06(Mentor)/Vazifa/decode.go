package main

import "fmt"

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key, message))
}
func decodeMessage(key string, message string) string {
	mapStr := make(map[rune]rune)
	var idx int32
	for _, v := range key {
		if _, ok := mapStr[v]; !ok && v != ' ' {
			mapStr[v] = 'a' + idx
			idx++
		}
	}

	res := []rune{}
	for _, v := range message {
		if v == ' ' {
			res = append(res, ' ')
		} else if _, ok := mapStr[v]; ok {
			res = append(res, mapStr[v])
		}
		return string(res)
	}
}
