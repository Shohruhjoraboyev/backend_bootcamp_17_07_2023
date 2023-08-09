package structmap

import "fmt"

// defer funksiyaning bajarilishini tashqi funksiya tugallanishigacha kechiktiradi
func deferF() {

	//1
	// defer bilan ishlatilgan funksiyalar argument qiymatlarini chaqirilgan paytni o'zida oladi
	name := "John"
	defer fmt.Println("Hello ", name)
	name = "Alex"
	fmt.Println("hello world")
	fmt.Println("name:", name)

	//2
	// defer funksiyalarni stackga soladi.
	//Tashqi funksiya tugallanganda defer qilingan funksiyalar
	//last-in-first-out tartibda bajariladi
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
