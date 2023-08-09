package structmap

import "fmt"

// map key:value juftligida ma'lumot saqlaydi
// zero value nil. nil mapda key bo'lmaydi va qo'shib ham bo'lmaydi
func mapFunc() {
	/*var m map[int]string
	if m == nil {
		fmt.Println("map is nil")
	}
	m[1] = "one"
	fmt.Println(m)

	anotherMap := make(map[int]string)
	anotherMap[1] = "one"
	fmt.Println(anotherMap)

	mapValued := map[string]string{
		"name":  "John Doe",
		"job":   "developer",
		"stack": "Golang",
	}

	fmt.Println(mapValued)*/
	//qiymat bilan e'lon qilish
	usersMap := map[string]user{
		"john": user{"John", 23},
		"alex": {Name: "Alex"},
	}
	fmt.Println(usersMap)

	// mapni elemetini olish
	//1
	elem := usersMap["john"]
	fmt.Println(elem)

	//2
	elem, ok := usersMap["eddie"]
	if ok {
		fmt.Println("topildi")
		fmt.Println(elem, ok)
	} else {
		fmt.Println("topilmadi")
		fmt.Println(elem, ok)
	}

	//element qo'shish
	usersMap["jack"] = user{"Jack", 34}
	fmt.Println(usersMap)

	//elementni o'chirish
	delete(usersMap, "jack")

}

type user struct {
	Name string
	Age  int
}
