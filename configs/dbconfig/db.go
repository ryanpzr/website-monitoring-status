package dbconfig

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Oscaveiras1! dbname=monitoring_db sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
