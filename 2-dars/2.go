package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func staircase(n int32) {
	for i := n - 1; i >= 0; i-- {
		fmt.Printf(printChar(i, " "))
		fmt.Printf(printChar(n-i, "#"))
		fmt.Printf("\n")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func printChar(x int32, s string) string {
	st := ""
	var i int32
	for i = 0; i < x; i++ {
		st += s
	}
	return st
}
