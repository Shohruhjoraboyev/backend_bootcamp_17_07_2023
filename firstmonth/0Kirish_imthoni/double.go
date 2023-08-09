package main

import "fmt"

func main() {
	unfiltered := []string{"a", "abc", "abc"}
	// unfiltered:=[]string{"bv","abc","abc"}
	// target:="abc"
	filterd := []string{}
	for _, v := range unfiltered {
		for _, char := range v {
			if string(char)==string(v) {
				filterd = append(filterd, string(char))

			}

		}

	}
	fmt.Println(filterd[0])

}
