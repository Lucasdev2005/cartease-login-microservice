package Core

import (
	"database/sql"
	"log"
)

var DatabaseConection *sql.DB

func Main() {
	conStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	DatabaseConection = db
}

func Query(qry *string) (*sql.Rows, error) {
	queryExecuted, err := DatabaseConection.Query(*qry)
	return queryExecuted, err
}
