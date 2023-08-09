package main

type Clients struct {
	Id   int
	Name string
	Card Card
}

type Card struct {
	Products []CardProducts
}

type CardProducts struct {
	ProductId int
	SizeId    int
	quantity  int
}

type Product struct {
	Id    int
	Name  string
	Sizes []Size
}

type Size struct {
	Id    int
	Name  string
	Price int
}

func main() {
	client := Clients{
		Id:   1,
		Name: "Shohruh",
		Card: Card{
			Products: []CardProducts{
				{
					ProductId: 1,
					SizeId:    1,
					quantity:  10,
				},
				{
					ProductId: 1,
					SizeId:    1,
					quantity:  10,
				},
			},
		},
	}
	client1 := Clients{
		Id:   2,
		Name: "Islom",
		Card: Card{
			Products: []CardProducts{
				{
					ProductId: 2,
					SizeId:    2,
					quantity:  3,
				},
				{
					ProductId: 1,
					SizeId:    1,
					quantity:  2,
				},
			},
		},
	}
	products := []Product{
		{
			Id:   1,
			Name: "Cola",
			Sizes: []Size{
				{
					Id:    1,
					Name:  "25sm",
					Price: 10000,
				},
				{
					Id:    2,
					Name:  "30sm",
					Price: 15000,
				},
				{
					Id:    3,
					Name:  "35sm",
					Price: 25000,
				},
			},
		},
		{
			Id:   2,
			Name: "Fanta",
			Sizes: []Size{
				{
					Id:    1,
					Name:  "25sm",
					Price: 5000,
				},
				{
					Id:    2,
					Name:  "30sm",
					Price: 1000,
				},
				{
					Id:    3,
					Name:  "35sm",
					Price: 15000,
				},
			},
		},
	}

}
