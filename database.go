package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
	ID        int    `json:"id"`
}

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)

	return db
}
