package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string  `json:"name"`
	Sentiment string  `json:"sentiment"`
	ID        []uint8 `json:"id"`
}

func EnsureDBIsOpen() *sql.DB {
	db := OpenDBConnection()
	return db
}

func OpenDBConnection() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_DB_NAME"), os.Getenv("RDS_HOSTNAME"), os.Getenv("RDS_PORT"))
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	// createTable, err := db.Prepare("CREATE TABLE hashtags (id serial primary key, hashtag varchar(180) unique, sentiment char(9));")
	// checkErr(err)

	res, err := db.Exec("CREATE TABLE hashtags (id serial primary key, hashtag varchar(180) unique, sentiment char(9));")
	checkErr(err)

	return db
}

// ensure db connection is open func
// store db object
// if no db obj, calls open + checks tables, creates if none exist

func ViewRows(db *sql.DB) []Hashtag {
	rows, err := db.Query("SELECT * FROM hashtags")
	checkErr(err)
	defer rows.Close()

	var hashtags []Hashtag

	for rows.Next() {
		var hashtag Hashtag
		err = rows.Scan(&hashtag.Name, &hashtag.Sentiment, &hashtag.ID)
		checkErr(err)
		hashtags = append(hashtags, hashtag)
	}

	return hashtags
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

func SelectRandomHashtag(db *sql.DB, sentiment string) Hashtag {
	var hashtag Hashtag
	err := db.QueryRow("SELECT * FROM hashtags WHERE sentiment = $1 ORDER BY random()", sentiment).Scan(&hashtag.ID, &hashtag.Name, &hashtag.Sentiment)
	checkErr(err)

	return hashtag
}
