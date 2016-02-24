package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InsertHashtag(db *sql.DB, hashtag string) {
	fmt.Println("# Inserting values")

	var lastInsertID int
	err := db.QueryRow("INSERT INTO hashtags(hashtag) VALUES($1) returning id;", hashtag).Scan(&lastInsertID)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertID)
}

func ShowAllHashtags(db *sql.DB) {
	fmt.Println(" -> Querying")
	rows, err := db.Query("SELECT * FROM hashtags")
	checkErr(err)

	fmt.Println(" hashtag | id ")

	for rows.Next() {
		var id int
		var hashtag string
		err = rows.Scan(&hashtag, &id)
		checkErr(err)
		fmt.Printf(" %v | %v \n", hashtag, id)
	}
}
