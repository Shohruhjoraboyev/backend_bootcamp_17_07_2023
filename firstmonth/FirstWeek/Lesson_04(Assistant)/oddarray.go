package main

import "fmt"

func main() {
	smth := [5]int{1, 2, 3, 4, 5}
	var numeven, evensum, oddnum, oddsum int
	evenarr := []int{}
	oddarr := []int{}
	for i := 0; i < len(smth); i++ {
		if smth[i]%2 == 0 {
			numeven++
			evensum += smth[i]
			evenarr = append(evenarr, smth[i])

		} else {
			oddnum++
			oddsum += smth[i]
			oddarr = append(oddarr, smth[i])

		}

	}
	fmt.Println("Count of Even numebers: ", numeven)
	fmt.Println("Sum of Even numebers: ", evensum)
	fmt.Println("Even numbers: ", evenarr)

	fmt.Println("Odd numbers: ", oddnum)
	fmt.Println("Sum of Odd numbers: ", oddsum)
	fmt.Println("Even numbers: ", oddarr)

}
