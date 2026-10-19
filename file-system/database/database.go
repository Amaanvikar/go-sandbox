package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() error {
	connStr := "postgres://amanpathan@localhost:5432/fileupload_db?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL successfully!")
	return nil
}
