package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ball int
	var score int
	fmt.Println("Tosh qaychi qog'oz o'yniga xush kelibsiz!!")
	fmt.Println("Necha achkogacham o'ynamoqchisz: ")
	fmt.Scan(&ball)
	fmt.Println("O'yin shartlari: ")
	fmt.Println("Tosh:1 ,Qaychi:2,Qog'oz:3")
	fmt.Println("Har yutuq 1 ball hisoblanadi")

	n := 4
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(n)
	var guess int
	var again string

	for ball > 0 && score <= ball {
		for try := 1; try <= ball; try++ {
			fmt.Println("Sizdagi bor ballar: ", ball)
			fmt.Println("Sizning jami ballaringiz: ", score)
			fmt.Println("Kiriting: ")
			fmt.Scan(&guess)
			if guess == 1 && secretNumber == 1 || guess == 2 && secretNumber == 22 || guess == 3 && secretNumber == 3 {
				fmt.Println("Durrang")
			} else if guess == 1 && secretNumber == 2 || guess == 2 && secretNumber == 3 || guess == 3 && secretNumber == 1 {
				fmt.Println("Barakalla sizga 1 ball")
				score++
			} else if guess == 1 && secretNumber == 3 || guess == 2 && secretNumber == 1 || guess == 3 && secretNumber == 2 {
				fmt.Println("Yutkazdingiz:")
			}

			if ball == n {
				fmt.Println("O'yin tugadi!!")
				time.Sleep(2 * time.Second)

				fmt.Printf("Tog'ri javob:  %d\n", secretNumber)
				return
			}

			fmt.Println("Yana o'ynamoqchi bo'lsangiz <ha> deb yozing")
			fmt.Scan(&again)
			if again != "ha" && again != "Ha" {
				fmt.Println("Ishtrokiz uchun raxmat")
				return
			}
		}
		ball--
	}

	fmt.Println("O'yin tugadi!!")
	time.Sleep(2 * time.Second)

	fmt.Printf("Tog'ri javob:  %d\n", secretNumber)
}
