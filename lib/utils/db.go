package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var databaseURL string = os.Getenv("DATABASE_URL")

func ConnectDBOrFatal() *sql.DB {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDBOrFatal(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func QueryDBOrFatal(db *sql.DB, query string, args ...any) *sql.Rows {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func QueryRowDBOrFatal(db *sql.DB, query string, args ...any) *sql.Row {
	row := db.QueryRow(query, args...)

	err := row.Err()
	if err != nil {
		log.Fatal(err)
	}

	return row
}

func CloseRowsOrFatal(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

type scanner interface {
	Scan(dest ...any) error
}

func ScanOrFatal(s scanner, dest ...any) {
	err := s.Scan(dest...)
	if err != nil {
		log.Fatal(err)
	}
}
