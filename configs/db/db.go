package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
