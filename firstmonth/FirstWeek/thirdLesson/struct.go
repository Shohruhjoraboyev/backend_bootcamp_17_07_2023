package structmap

import "fmt"

type Person struct {
	Name string
	Age  int
	Parent
}
type Parent struct {
	Name string
	Age  int
	a    []string
}

//struct bir qancha turli tipdagi qiymatlarni saqlaydi
//structda qiymatlar uning fieldlarida saqlanadi
func structF() {
	//1
	fmt.Println(Person{Name: "John"})
	//field_namelarni ko'rsatmaslik ham mumkin,
	//bunda beriladigan qiymatlar struct e'lon qilingan tartibda berilishi kerak
	fmt.Println(Person{"John", 23, Parent{"Jack", 45}}) //hamma fieldlarga qiymat berish majburiy

	//2
	var (
		john Person
		alex = Person{}
	)
	fmt.Println(john)
	fmt.Println(alex)

	john = Person{"John", 23}
	fmt.Println(john)

	// struct fieldi struct.field_name ko'rinishida ishlatiladi
	alex.Name = "Alex"
	alex.Age = 65
	fmt.Println(alex)

}
