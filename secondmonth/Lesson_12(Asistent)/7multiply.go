/*
Exercise 7: Create a program that uses goroutines
and channels to multiply each element in a given slice by 2 concurrently.
*/
package main

import "fmt"

func main() {
	numbers := []int{2, 4, 5, 3}
	ch := make(chan []int)
	go multiply(numbers, ch)
	fmt.Println(<-ch)
}

func multiply(numbers []int, ch chan []int) {
	slice := []int{}

	for _, v := range numbers {
		m := v * 2
		slice = append(slice, m)
	}

	ch <- slice
	close(ch)
}

SELECT c.name, MAX(s.count) AS TopSelled
FROM company AS c
JOIN sale AS s ON c.id = s.company_id
GROUP BY c.name
order by  TopSelled;