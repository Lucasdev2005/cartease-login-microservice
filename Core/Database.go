package Core

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DatabaseConnection *sql.DB

func init() {
	conStr := "user=cartease dbname=db password=cartease@123 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConnection = db
}

func Query(qry string) (*sql.Rows, error) {
	queryExecuted, err := DatabaseConnection.Query(qry)
	return queryExecuted, err
}
