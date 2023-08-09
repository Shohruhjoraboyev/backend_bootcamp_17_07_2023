package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Println("Qaysi oraliqgacha sonlardan topmoqchisiz: ")
	fmt.Scan(&n)
	tryNum := int(math.Log2(float64(n)))

	fmt.Println("Urinishlar soni: ", tryNum)
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100)
	fmt.Println("Secret", secretNumber)

	var guess int
	var again string
	for true {
		for try := 1; try <= tryNum; try++ {

			fmt.Printf("Imkoniyat:  %d\n", try)

			fmt.Println("Iltimos o'z taxminizni kiriting")
			fmt.Scan(&guess)
			if guess > n {
				fmt.Println("Siz oraliqdagi sondan katta qiymat kiritdiz: ")
				fmt.Printf("Iltmos shu %d oralig'idagi sonni kiriting: ", n)
			} else if guess < secretNumber {
				fmt.Printf("Sizni  taxminingiz xato,biroz kattaroq son o'ylab ko'ring\n")
			} else if guess > secretNumber {
				fmt.Printf("Sizni taxminingiz xato: biroz kichikroq  son o'ylab toping\n")
			} else {
				fmt.Printf("Siz to'g'ri topdingiz\n")
				break
			}
			if try > tryNum {
				// 5ta imkoniyatdan kegin oyin tugaydi va tepadagi tsikl tugaydi
				fmt.Println("O'yin tugadi!!")
				time.Sleep(2 * time.Second)

				fmt.Printf("Tog'ri javob:  %d\n", secretNumber)
				break
			}
		}
		fmt.Println("Yana o'ynamoqchi bo'lsangiz <ha> deb yozing")
		fmt.Scan(&again)
		if again != "ha" || again != "Ha" {
			fmt.Println("Ishtrokiz uchun raxmat")
			break
		}
	}
}
