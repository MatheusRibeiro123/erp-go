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
// Função para obter um cliente por ID *ainda não esta em funcionamento, falta corrigir

func (r *ClientRepository) GetByID(id int) (models.Client, error) {
    query := "SELECT id , name , email , phone , document , created_at FROM clients WHERE id = $1"

	row := r.DB.QueryRow(query, id)
	
	var client models.Client

	err = row.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Document, &client.CreatedAt)
	
	if err != nil {
		return nil, err
	}
	return client, nil
}