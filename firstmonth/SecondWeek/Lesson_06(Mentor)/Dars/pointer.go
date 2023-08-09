package structmap

// Pointer qiymat joylashgan xotirani addressini saqlaydi
// *Type bu Typedagi qiymatga pointer. Pointer uchun zero value nil
func main() {
	// //1
	// name, age := "John", 19
	// var namePointer *string

	// // & operator variable bilan ishlatiladi va uning qiymatini addressini oladi
	// namePointer = &name
	// agePointer := &age

	// fmt.Println(namePointer, agePointer)

	// //2
	// // * operator ponter bilan ishlatiladi va undagi qiymatni olib beradi
	// fmt.Println(*namePointer, *agePointer)

	// //3
	// //pointer orqali addressdagi qiymatni o'zgartirish
	// *namePointer = "Alex"
	// *agePointer = *agePointer + 5

	// fmt.Println(name, age)
	pointerToStruct()
}

// pointer to structs

type Car struct {
	Model       string
	ReleaseDate string
}

func pointerToStruct() {
	var spark = Car{"Spark", "2021-10-30"}

	p := &spark
	p.ReleaseDate = "2021-11-25"
	println(spark)
}
