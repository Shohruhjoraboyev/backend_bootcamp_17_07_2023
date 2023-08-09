//cmd faylda  - main.go boshqa packegelar main functionda init qilinadi.( length bn slice ochish .env ni o'qish)
//config -- iota  ,static varaiblelar saqlanadi, .env dan o'qiladgon variablellarni o'qish
//handler-- programmaga request kelganda handle qiladgon  funksiya yoziladi
//Models--programma ishlatilinadgon modellar  "bunda logika bo'lmedi"
//storage-- malumotlar saqlanadgon joy bu database bo'lishi mumkin.

package main

//Interface metodlar toplami
//kamroq kod yozish va tushunarlikroq bo'lishiga yordam beradi.
//barcha methodlarini bajaradigon tiplarini foydalana oladi  interface xussusiyati
//bosh interface ag xoxlgan tipdagi qiymatni olsa bo'ladi
//bosh interfaceda method bo'lmaydi

import "fmt"

type Stringer interface {
	String() string
}
type Book struct {
	Name string
}

type Journal struct {
	Title string
}
type myInt int

func (i myInt) String() string {
	return fmt.Sprintf("%d", i)
}

func (b Book) String() string {
	return b.Name
}

func (j Journal) String() string {
	return j.Title
}
func main() {
	var b, j, i Stringer
	b = Book{Name: "O'tkan kunlar"}
	j = Journal{"Yosh kuch"}
	i = myInt(5)
	fmt.Println(b.String())
	fmt.Println(j.String())
	fmt.Println(i.String())
}
