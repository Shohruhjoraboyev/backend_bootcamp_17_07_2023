package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Println("Please enter radius : ")
	fmt.Scanln(&r)
	L := float32(2) * 3.14 * float32(r)
	s := 3.14 * float32(r) * float32(r)

	fmt.Println("Kvadrat uzunligi: ", L)
	fmt.Println("Kvadrat yuzasi: ", s)
}
