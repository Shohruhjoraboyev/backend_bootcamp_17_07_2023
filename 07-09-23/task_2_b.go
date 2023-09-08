package main

// import (
// 	"fmt"
// 	"strings"
// 	"sync"
// )

// func main() {
// 	word := "misol"
// 	sentences := []string{
// 		"Bu birinchi misol so'zdan iborat",
// 		"Ikkinchi misolga kiring, shu so'zni qidiring",
// 		"Uchinchi misolning oxiriga yetib kelganimizda, so'zni topib chiqaring",
// 	}

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	var foundSentences []string

// 	wg.Add(len(sentences))

// 	for _, sentence := range sentences {
// 		go func(s string) {
// 			defer wg.Done()
// 			if strings.Contains(s, word) {
// 				mu.Lock()
// 				foundSentences = append(foundSentences, s)
// 				mu.Unlock()
// 			}
// 		}(sentence)
// 	}

// 	wg.Wait()

// 	fmt.Println("So'z qatnashgan gaplar:")
// 	for _, found := range foundSentences {
// 		fmt.Println(found)
// 	}
// }
