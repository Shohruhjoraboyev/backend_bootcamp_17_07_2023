package main

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// )

// func countLines(filename string, wg *sync.WaitGroup, ch chan<- int) {
// 	defer wg.Done()

// 	content, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Printf("Error reading file %s: %v\n", filename, err)
// 		ch <- 0
// 		return
// 	}

// 	lines := 0
// 	for _, char := range content {
// 		if char == '\n' {
// 			lines++
// 		}
// 	}

// 	ch <- lines
// }

// func main() {
// 	files := []string{"1.go", "2.go", "3.go"}
// 	mapFiles := make(map[string]int)

// 	var wg sync.WaitGroup
// 	ch := make(chan int, len(files))

// 	for _, file := range files {
// 		wg.Add(1)
// 		go countLines(file, &wg, ch)
// 		mapFiles[file] = <-ch
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for file, line := range mapFiles {
// 		fmt.Printf("%s -> number of lines: %d\n", file, line)
// 	}
// }
