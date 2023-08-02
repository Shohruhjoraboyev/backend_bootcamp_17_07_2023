package main

type Client struct {
	ID   int
	Name string
	Card Card
}

type Card struct {
	Products []CardProduct
}

type CardProduct struct {
	ProductID int
	SizeID    int
	Quantity  int
}

type Product struct {
	ID    int
	Name  string
	Sizes []Size
}

type Size struct {
	ID    int
	Name  string
	Price int
}

func main() {

}
