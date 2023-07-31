package cruds

type Client struct {
	ID   int
	Name string
	Card Card
}

type Card struct {
	Products []CardProducts
}

type CardProducts struct {
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
	Price float64
}
