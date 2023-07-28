package main

import "fmt"

func main() {
	a := 460

	if a <= 90 {
		fmt.Println("1- chorak")
	} else if a > 90 && a <= 180 {
		fmt.Println("2-chorak")
	} else if a > 180 && a <= 270 {
		fmt.Println("3-chorak")
	} else if a > 270 && a <= 360 {
		fmt.Println("4-chorak")
	} else if a > 360 {
		a = a - 360
		if a <= 90 {
			fmt.Println("1- chorak")
		} else if a > 90 && a <= 180 {
			fmt.Println("2-chorak")
		} else if a > 180 && a <= 270 {
			fmt.Println("3-chorak")
		} else if a > 270 && a <= 360 {
			fmt.Println("4-chorak")
		}
	}
}
