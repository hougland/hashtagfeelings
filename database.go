package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
)

// this func is just for practice with databases and Go
func ShowAllHashtags(db *sql.DB) {
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM hashtags")
	checkErr(err)
	defer rows.Close()

	fmt.Println(" hashtag | id ")

	for rows.Next() {
		var id int
		var hashtag string
		err = rows.Scan(&hashtag, &id)
		checkErr(err)
		fmt.Printf(" %v | %v \n", hashtag, id)
	}
}

func SelectOneHashtag(db *sql.DB) {
	// does the randomly generating numbers thing actually work?
	// needs to be altered to return the hashtag

	fmt.Println("# Selecting Random Hashtag")

	var numRows int
	err := db.QueryRow("SELECT count(*) FROM hashtags").Scan(&numRows)
	checkErr(err)

	randInt := rand.Intn(numRows)

	var hashtag string
	err = db.QueryRow("SELECT hashtag FROM hashtags where id = $1", randInt).Scan(&hashtag)
	checkErr(err)

	fmt.Println(hashtag)
}

func SaveSentiment() {
	// save sentiment object in database
	// will be similar to InsertHashtag func
}

func InsertHashtag(db *sql.DB, hashtag string) {
	fmt.Println("# Inserting values")

	var lastInsertID int
	err := db.QueryRow("INSERT INTO hashtags(hashtag) VALUES($1) returning id;", hashtag).Scan(&lastInsertID)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertID)
}
