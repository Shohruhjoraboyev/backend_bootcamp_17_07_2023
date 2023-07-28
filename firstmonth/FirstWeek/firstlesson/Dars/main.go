package main

import "fmt"

func main() {
	//ozgaruvchilarni qiymatini funksiya orqalik ham bersa bo'ladi
	// global o'zgaruvchilarni elon qilishda : nuqta bn  elon qilib bo'lmaydi.
	// : bilan yaratilgan o'zgaruvchi yaratilgan  scoopni ichda ishlaydi.
	// interface ham data type unda struct va methodlar bor structlar o'zgaruvchilarni qabul qilish uchun methodlar esa shu o'zgaruvchilarda funksya bajarish uchun
	//bo'sh  interface degan atama ham bor u methodga keladgon variable ni data type aniq bo'lmagan holatda ishlatiladi.
	var name, age = nameAge()
	sekai := "世界中"
	fmt.Printf("The your value  is: %s\n", name)
	fmt.Printf("The your value  is: %d\n", age)
	fmt.Printf("Sekaijyuu: %s", sekai)
}
func nameAge() (string, int) {
	return "John", 16
}
