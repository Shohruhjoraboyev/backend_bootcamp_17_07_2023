package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
// }

// func decodeMessage(key string, message string) string {
// 	mapString := make(map[rune]rune)
// 	idx := 0
// 	for _, v := range key {
// 		if v == ' ' {
// 			continue
// 		}

// 		if _, ok := mapString[v]; !ok {
// 			mapString[v] = rune(idx) + 'a'
// 			idx += 1
// 		}
// 		// if idx == 26 {
// 		// 	break
// 		// }
// 	}

// 	res := []rune{}
// 	for _, v := range message {
// 		if v == ' ' {
// 			res = append(res, ' ')
// 		} else {
// 			res = append(res, mapString[v])
// 		}
// 	}

// 	return string(res)
// }
