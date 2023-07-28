package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Bu birinchi ishlashim kerak bo'lgan masala edi.uni va nixoyat yechimini  va topdim"
	words := strings.Fields(text)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++

	}

	for word, count := range counts {
		fmt.Printf("%s: %d marta\n", word, count)
	}

}
