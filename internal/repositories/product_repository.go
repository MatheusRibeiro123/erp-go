package repositories

import (
	"database/sql"
	"erp-go/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

// função para obter todos os produtos do banco de dados
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

// função para obter um produto específico pelo ID
func (r *ProductRepository) GetByID(id int) (models.Product, error) {
	query := "SELECT id, name, description, stock_quantity, price, created_at FROM products WHERE id = $1"

	row := r.DB.QueryRow(query, id)

	var product models.Product

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.StockQuantity, &product.Price, &product.CreatedAt)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

// função para criar um novo produto no banco de dados
func (r *ProductRepository) Create(product models.Product) (int, error) {
	query := "INSERT INTO products (name, description, stock_quantity, price) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int

	err := r.DB.QueryRow(query, product.Name, product.Description, product.StockQuantity, product.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// função para atualizar um produto existente no banco de dados
func (r *ProductRepository) Update(id int, product models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, stock_quantity = $3, price = $4 WHERE id = $5"

	_, err := r.DB.Exec(query, product.Name, product.Description, product.StockQuantity, product.Price, id)
	if err != nil {
		return err
	}

	return nil
}

// função para excluir um produto do banco de dados
func (r *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
