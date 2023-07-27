package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	sozlar := strings.Fields(text)

	res := make(map[string]int)
	for _, v := range sozlar {
		res[v] += 1
	}

	fmt.Println(res)
}
