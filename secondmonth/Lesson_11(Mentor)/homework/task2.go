package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "cat"
	sentences := []string{
		"My neighbor has a cute black cat named Midnight.",
		"I enjoy watching funny cat videos on the internet.",
		"The cat curled up on the windowsill and basked in the sunlight.",
		"She adopted a stray cat and gave it a loving home.",
	}
	ch := make(chan string)
	gapCh := make(chan []string)
	go GetOneSentences(word, sentences, ch, gapCh)
	//Case A
	fmt.Println("Sentences containing word:", <-ch)
	// Case B
	fmt.Println("All sentences containing word:", <-gapCh)
}

func GetOneSentences(word string, sentences []string, ch chan string, gapCh chan []string) {
	result := ""
	gap := []string{}
	for _, sentence := range sentences {
		if strings.Contains(sentence, word) {
			result = sentence
			gap = append(gap, sentence)
		}
	}
	ch <- result
	gapCh <- gap
	close(ch)
	close(gapCh)
}
