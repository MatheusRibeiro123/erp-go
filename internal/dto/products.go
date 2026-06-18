package dto

// Product DTOs
type CreateProductInput struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

type UpdateProductInput struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

type PatchProductInput struct {
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	StockQuantity *int     `json:"stock_quantity"`
}
