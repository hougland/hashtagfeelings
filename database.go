package main

import (
	"database/sql"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
	ID        int    `json:"id"`
}

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=BluePenguin dbname=ebdb sslmode=disable")

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

// func InsertHashtag(db *sql.DB, hashtag string, sentiment string) {
// 	stmt, err := db.Prepare("INSERT INTO hashtags(hashtag, sentiment) VALUES($1, $2);")
// 	checkErr(err)
// 	defer stmt.Close()
//
// 	_, err = stmt.Exec(hashtag, sentiment)
// 	checkErr(err)
// }

// func SelectRandomHashtag(db *sql.DB, sentiment string) Hashtag {
// 	var hashtag Hashtag
// 	err := db.QueryRow("SELECT * FROM hashtags WHERE sentiment = $1 ORDER BY random()", sentiment).Scan(&hashtag.ID, &hashtag.Name, &hashtag.Sentiment)
// 	checkErr(err)
//
// 	return hashtag
// }
