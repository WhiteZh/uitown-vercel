package utils

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var databaseURL string = os.Getenv("DATABASE_URL")

func TryConnectDB(w http.ResponseWriter) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		WriteInternalErrorResponse(w)
		log.Fatal(err)
	}

	return db
}

func TryCloseDB(db *sql.DB, w http.ResponseWriter) {
	if err := db.Close(); err != nil {
		WriteInternalErrorResponse(w)
		log.Fatal(err)
	}
}
