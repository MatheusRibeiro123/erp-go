package repositories

import (
	"database/sql"
	"erp-go/internal/apperrors"
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
		return nil, apperrors.TranslatePostgresError(err)
	}

	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.StockQuantity, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.TranslatePostgresError(err)
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
		if err == sql.ErrNoRows {
			return models.Product{}, apperrors.ErrNotFound
		}
		return models.Product{}, apperrors.TranslatePostgresError(err)
	}

	return product, nil

}

// função para criar um novo produto no banco de dados
func (r *ProductRepository) Create(product models.Product) (int, error) {
	query := "INSERT INTO products (name, description, stock_quantity, price) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int

	err := r.DB.QueryRow(query, product.Name, product.Description, product.StockQuantity, product.Price).Scan(&id)
	if err != nil {
		return 0, apperrors.TranslatePostgresError(err)
	}

	return id, nil
}

// função para atualizar um produto existente no banco de dados
func (r *ProductRepository) Update(id int, product models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, stock_quantity = $3, price = $4 WHERE id = $5"

	result, err := r.DB.Exec(query, product.Name, product.Description, product.StockQuantity, product.Price, id)
	if err != nil {
		return apperrors.TranslatePostgresError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

// função para excluir um produto do banco de dados
func (r *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return apperrors.TranslatePostgresError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}
