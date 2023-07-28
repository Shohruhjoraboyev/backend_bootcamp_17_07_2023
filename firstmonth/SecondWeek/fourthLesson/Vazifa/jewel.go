package main

import (
	"fmt"
)

func main() {
	jewels := "aA"
	stones := "aAAbbBB"
	fmt.Println(jewel(jewels, stones))
}

func jewel(jewels string, stones string) int {
	count := 0
	for _, j := range jewels {
		for _, s := range stones {
			if j == s {
				count++
			}
		}
	}
	return count
}
