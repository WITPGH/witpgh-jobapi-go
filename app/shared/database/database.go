package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var WITPGH *sql.DB

func ConnectWITJobBoard() {
	var err error

	// Connect to PostgreSQL
	if WITPGH, err = sql.Open("postgres", getConnection()); err != nil {
		log.Println("SQL Driver Error", err)
	}

	// Check if is alive
	if err = WITPGH.Ping(); err != nil {
		log.Println("WIT Job Board Database Error", err)
	}
}

func getConnection() string {
	return os.Getenv("DATABASE_URL")
}
