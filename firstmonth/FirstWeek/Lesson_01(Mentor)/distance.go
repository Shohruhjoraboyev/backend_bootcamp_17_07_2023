package main

import (
	"fmt"
	"math"
)

func main() {
	x1, x2, y1, y2 := 2.0, 1.0, 2.0, 1.0
	dx := x2 - x1
	dy := y2 - y1
	fmt.Println("Distance between of 2 point is: ", math.Sqrt(dx*dx+dy*dy))
}
