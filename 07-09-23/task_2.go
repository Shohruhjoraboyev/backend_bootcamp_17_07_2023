package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	word := "misol"
	sentences := []string{
		"Bu birinchi misol so'zdan iborat",
		"Ikkinchi misolga kiring, shu so'zni qidiring",
		"Uchinchi misolning oxiriga yetib kelganimizda, so'zni topib chiqaring",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var foundSentence string

	for _, sentence := range sentences {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			if strings.Contains(s, word) {
				mu.Lock()
				if foundSentence == "" {
					foundSentence = s
				}
				mu.Unlock()
			}
		}(sentence)
	}

	wg.Wait()

	if foundSentence != "" {
		fmt.Printf("Topilgan gap: %s\n", foundSentence)
	} else {
		fmt.Println("Topilmadi")
	}
}
