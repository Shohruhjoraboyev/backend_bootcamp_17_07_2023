package main

import "fmt"

func main() {

	var a = []int{1, 2, 3}
	var b = []int{3, 2, 1}
	var c = []int{4, 4, 3, 1, 4}

	fmt.Println(compareTriplets(a, b))

	staircase(6)

	fmt.Println(birthdayCakeCandles(c))

	fmt.Println(reverseNumber(123))

}

func compareTriplets(a []int, b []int) []int {

	var counterA, counterB int

	for i := 0; i < len(a); i++ {

		if a[i] < b[i] {
			counterB++
		} else if a[i] > b[i] {
			counterA++
		}

	}

	var result = [2]int{counterA, counterB}

	return result[:]

}

func staircase(n int) {

	for i := 1; i <= int(n); i++ {

		var result string

		for j := i; j < int(n); j++ {
			result += " "
		}

		for k := 1; k <= i; k++ {

			result += "#"

		}

		fmt.Println(result)

	}

}

func birthdayCakeCandles(candles []int) int {

	var counter int
	var max = candles[0]

	for i := 0; i < len(candles); i++ {

		if max <= candles[i] {

			max = candles[i]

		}

	}

	for i := 0; i < len(candles); i++ {

		if max == candles[i] {
			counter++
		}

	}

	return counter

}

func reverseNumber(num int) int {

	var dig, result int

	for num != 0 {

		dig = num % 10
		result = result*10 + dig

		num = (num - dig) / 10

	}

	return result

}
