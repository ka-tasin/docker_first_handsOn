package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orane",
		Description: "Best orange ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Best apple ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Best banana ever",
		Price:       20,
		ImageUrl:    "https://www.quanta.org/orange/",
	}

	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
}
