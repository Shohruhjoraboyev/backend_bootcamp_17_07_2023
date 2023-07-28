package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Example 1:")
	// var str string = "Live young live free"
	// var substr string = "young"
	// if strings.Contains(str, substr) {
	// 	fmt.Printf("The string '%s' contains the substring '%s'.\n", str, substr)
	// } else {
	// 	fmt.Printf("The string '%s' does not contain the substring '%s'.\n", str, substr)
	// }
	str := "Salom"
	substr := "al"

	if strings.Contains(str, substr) {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}

}
