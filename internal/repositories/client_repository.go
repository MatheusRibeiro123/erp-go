package repositories

import (
	"database/sql"
)	

type ClientRepository struct {
	DB *sql.DB
}

func(r *ClientRepository) GetAll() {
	SELECT id , name , email , phone , document , created_at FROM clients

	rows, err := r.DB.Query()
}