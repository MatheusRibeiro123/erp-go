package repositories

import (
	"database/sql"
	"erp-go/internal/apperrors"
	"erp-go/internal/models"
	"errors"
)

type ClientRepository struct {
	DB *sql.DB
}

var ErrClientNotFound = errors.New("client not found")

// função para obter todos os clientes do banco de dados

func (r *ClientRepository) GetAll() ([]models.Client, error) {
	query := "SELECT id , name , email , phone , document , created_at FROM clients"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, apperrors.TranslatePostgresError(err)
	}
	defer rows.Close()

	var clients []models.Client

	for rows.Next() {

		var client models.Client

		err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Document, &client.CreatedAt)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.TranslatePostgresError(err)
	}
	return clients, nil

}

// função para obter um cliente específico pelo ID

func (r *ClientRepository) GetByID(id int) (models.Client, error) {
	query := "SELECT id , name , email , phone , document , created_at FROM clients WHERE id = $1"

	row := r.DB.QueryRow(query, id)

	var client models.Client

	err := row.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Document, &client.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Client{}, apperrors.ErrNotFound
		}
		return models.Client{}, apperrors.TranslatePostgresError(err)
	}
	return client, nil
}

//função para criar um novo cliente no banco de dados

func (r *ClientRepository) Create(client models.Client) (int, error) {
	query := "INSERT INTO clients (name, email, phone, document) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int

	err := r.DB.QueryRow(query, client.Name, client.Email, client.Phone, client.Document).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// função para atualizar um cliente existente no banco de dados

func (r *ClientRepository) Update(id int, client models.Client) error {

	query := "UPDATE clients SET name = $1, email = $2, phone = $3, document = $4 WHERE id = $5"

	result, err := r.DB.Exec(query, client.Name, client.Email, client.Phone, client.Document, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrClientNotFound
	}

	return nil

}

// função para excluir um cliente do banco de dados

func (r *ClientRepository) Delete(id int) error {

	query := "DELETE FROM clients WHERE id = $1"

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrClientNotFound
	}

	return nil
}
