package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("task-1")
	var son int
	fmt.Print("Sonni kiriting: ")
	fmt.Scanf("%d", &son)

	if son%2 == 1 {
		fmt.Printf("%d - toq son.\n", son)
	} else {
		fmt.Printf("%d - juft son.\n", son)
	}

	fmt.Println("task-2")
	var son1, son2 int
	fmt.Print("1-son: ")
	fmt.Scan(&son1)
	fmt.Print("2-son: ")
	fmt.Scan(&son2)

	son1 = son1 + son2
	son2 = son1 - son2
	son1 = son1 - son2
	fmt.Printf("1-son: %d\n2-son: %d\n", son1, son2)

	fmt.Println("task-3")
	var num1, num2, num3 = 19, 12, 23
	fmt.Println("numbers: 19,12,23")
	var max, min int
	// find max number
	max = num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}
	// find min number
	min = num1
	if num2 < min {
		min = num2
	}
	if num3 < min {
		min = num3
	}
	fmt.Println("max + min =", min+max)

	fmt.Println("task-4")
	var x1, x2, y1, y2, dis float64

	x1, y1 = 2, 7
	x2, y2 = 5, 8

	fmt.Printf("x1,y1 = %.2f, %.2f\n", x1, y1)
	fmt.Printf("x2,y2 = %.2f, %.2f\n", x2, y2)

	dis = math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	fmt.Printf("distance: %.2f\n", dis)

	fmt.Println("task-5")
	var a, b, c float64
	fmt.Print("a,b,c ni kiriting, (misol: 1 2 3): ")
	fmt.Scan(&a, &b, &c)
	D := b*b - 4*a*c
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / 2 / a
		x2 := (-b - math.Sqrt(D)) / 2 / a
		fmt.Printf("x1 = %.2f\nx2 = %.2f\n", x1, x2)
	} else if D == 0 {
		x1 = -b / 2 / a
		fmt.Printf("x = %.2f\n", x1)
	} else {
		fmt.Println("haqiqiy yechimi yo'q")
	}
}
