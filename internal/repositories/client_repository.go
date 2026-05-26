package repositories

import (
	"database/sql"
	"erp-go/internal/models"
)	

type ClientRepository struct {
	DB *sql.DB
}

func (r *ClientRepository) GetAll() ([]models.Client, error) {
	query := "SELECT id , name , email , phone , document , created_at FROM clients"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return clients, nil

}


func (r *ClientRepository) GetByID(id int) (models.Client, error) {
    query := "SELECT id , name , email , phone , document , created_at FROM clients WHERE id = $1"

	row := r.DB.QueryRow(query, id)
	
	var client models.Client

	err := row.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Document, &client.CreatedAt)
	
	if err != nil {
		return models.Client{}, err
	}
	return client, nil
}

func (r *ClientRepository) Create(client models.Client) (int, error) {
	query := "INSERT INTO clients (name, email, phone, document) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int

	err := r.DB.QueryRow(query, client.Name, client.Email, client.Phone, client.Document).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}