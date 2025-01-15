package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var databaseURL string = os.Getenv("DATABASE_URL")

func ConnectDBOrPanic() *sql.DB {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Panic(err)
	}

	return db
}

func CloseDBOrPanic(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Panic(err)
	}
}

func ExecDBOrPanic(db *sql.DB, query string, args ...any) {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Panic(err)
	}
}

func QueryDBOrPanic(db *sql.DB, query string, args ...any) *sql.Rows {
	rows, err := db.Query(query, args...)

	if err != nil {
		log.Panic(err)
	}

	return rows
}

func QueryRowDBOrPanic(db *sql.DB, query string, args ...any) *sql.Row {
	row := db.QueryRow(query, args...)

	err := row.Err()
	if err != nil {
		log.Panic(err)
	}

	return row
}

func CloseRowsOrPanic(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Panic(err)
	}
}

type scanner interface {
	Scan(dest ...any) error
}

func ScanOrPanic(s scanner, dest ...any) {
	err := s.Scan(dest...)
	if err != nil {
		log.Panic(err)
	}
}
