package entity

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
}

type ProductWithCategory struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Price        float64  `json:"price"`
	CategoryID   int      `json:"category_id"`
	CategoryName string   `json:"category_name"`
}

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
}
