package main

import (
	"database/sql"
	"os"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string  `json:"name"`
	Sentiment string  `json:"sentiment"`
	ID        []uint8 `json:"id"`
}

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)

	return db
}

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
