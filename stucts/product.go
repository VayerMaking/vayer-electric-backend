package stucts

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SubcategoryId int64   `json:"subcategory"`
	Price       float64 `json:"price"`
	CurrentInventory   int64   `json:"inventory"`
	Image       string  `json:"image"`
	Brand       string  `json:"brand"`
	Sku         string  `json:"sku"`
	CreatedAt   string  `json:"created_at"`
}