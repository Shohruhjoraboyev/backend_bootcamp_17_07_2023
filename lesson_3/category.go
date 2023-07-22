package main

func main() {

}

// 3. Hamma ma'lumot bor, endi har bitta userni qaysi categorydan nechta
// favourite producti borligini topish kerak(loop,map,function,pointer .... ishlatib)

type Product struct {
	Id    int
	name  string
	price string
}

type Category struct {
	Id       int
	Name     string
	Products map[int]Product
}

type User struct {
	Id                    int
	Name                  string
	favourite_products_id []int
}

var ProductMap = map[int]Product{
	1: {
		Id:    1,
		name:  "apple",
		price: "$1",
	},
	2: {
		Id:    2,
		name:  "melon",
		price: "$3",
	},
	3: {
		Id:    3,
		name:  "laptop",
		price: "$2000",
	},
	4: {
		Id:    4,
		name:  "PC",
		price: "$2300",
	},
	5: {
		Id:    5,
		name:  "shirt",
		price: "$1",
	},
	6: {
		Id:    6,
		name:  "t-shirt",
		price: "$1",
	},
	7: {
		Id:    7,
		name:  "sock",
		price: "$1",
	},
	8: {
		Id:    8,
		name:  "air-cooler",
		price: "$1",
	},
	9: {
		Id:    9,
		name:  "banana",
		price: "$1",
	},
}

var UserMap = map[int]User{
	1: {
		Id:                    1,
		Name:                  "Omadbek",
		favourite_products_id: []int{1, 2, 3},
	},
}
