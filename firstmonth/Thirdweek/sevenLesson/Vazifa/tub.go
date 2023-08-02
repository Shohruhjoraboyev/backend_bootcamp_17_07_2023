//Berilgan n sonigacha barcha tub sonlarni chiqaring. Tub son - faqat o'ziga va 1ga bo'linadigan son(2,3,5,7.......)
package main

import "fmt"

func main() {
	a := primeNumbers()
	fmt.Println(a(7))
}

func primeNumbers() func(int) []int {
	return func(x int) []int {
		primes := []int{}
		for i := 0; i <= x; i++ {
			if i%2 != 0 {
				primes = append(primes, i)
			}

		}
		return primes
	}
}
