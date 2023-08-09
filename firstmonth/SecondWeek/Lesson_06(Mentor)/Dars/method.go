package structmap

//1
//Goda klasslar yo'q. Lekin typelarga methodlarni ulash mumkin
//Method receiver argument qabul qiluvchi funksiya
type User struct {
	Firstname, Lastname string
	Age                 int
}

//receiver func keyword va method_name orasiga yoziladi
//value receiver
func (U User) Fullname() (string, string) {
	return U.Firstname, U.Lastname
}

func FullnameFunc(absd User) string {
	return absd.Firstname + " " + absd.Lastname
}

//2
//structdan boshqa tiplar bilan ham method ishlatsa bo'ladi
type myFloat float64

func (f myFloat) Abs() myFloat {
	if f < 0 {
		return -f
	} else {
		return f
	}
}

//3
//pointerga method ulash
//bu usul amaliyotda ko'p ishlatiladi
//bunda funksiya ichida receiverning fieldlari
//qiymati o'zgarsa xotiradagi qiymat ham o'zgaradi
//pointer receiver
func (u *User) ChangeFirstname(newName string) {
	u.Firstname = newName
	return
}

//value receiverli methodlarda function valueni copy qilib ishlatadi
