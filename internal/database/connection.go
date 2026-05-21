package database

import ("database/sql"
	"log"
	_"github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	// TODO: move connection string to environment variables
	connStr := "future connection string"

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	log.Println("Connected to PostgreSQL!")
}