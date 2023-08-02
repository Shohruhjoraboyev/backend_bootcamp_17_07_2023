package main

import "fmt"

func main() {
	sport := struct {
		Id   int
		Name string
	}{Id: 1, Name: "football"}
	fmt.Println(sport)
}
