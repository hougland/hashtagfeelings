package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

func OpenDBConnection() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "BluePenguin", "hashtagfeelings")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	return db
}

func IsInTable(db *sql.DB, trend anaconda.Trend) bool {
	var id int
	err := db.QueryRow("SELECT id FROM hashtags WHERE hashtag = $1", trend.Name).Scan(&id)
	if err == nil {
		return true
	} else if err == sql.ErrNoRows {
		return false
	} else {
		panic(err)
	}
}

func InsertHashtag(db *sql.DB, hashtag string, sentiment string) {
	stmt, err := db.Prepare("INSERT INTO hashtags(hashtag, sentiment) VALUES($1, $2);")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(hashtag, sentiment)
	checkErr(err)
}

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

func SelectOneHashtag(db *sql.DB) string {
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

	return hashtag
}
