package repositories

import (
	"database/sql"
	"erp-go/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {

	query := "SELECT id, name, description, stock_quantity, price, created_at FROM products"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.StockQuantity, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
