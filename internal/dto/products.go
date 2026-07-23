package dto

// Product DTOs
type CreateProductInput struct {
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	Price         *float64 `json:"price" binding:"required,gte=0"`
	StockQuantity *int     `json:"stock_quantity" binding:"required,gte=0"`
}

type UpdateProductInput struct {
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	Price         *float64 `json:"price" binding:"required,gte=0"`
	StockQuantity *int     `json:"stock_quantity" binding:"required,gte=0"`
}

type PatchProductInput struct {
	Name          *string  `json:"name" binding:"omitempty,min=1"`
	Description   *string  `json:"description" binding:"omitempty,min=1"`
	Price         *float64 `json:"price" binding:"omitempty,gte=0"`
	StockQuantity *int     `json:"stock_quantity" binding:"omitempty,gte=0"`
}