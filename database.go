package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string  `json:"name"`
	Sentiment string  `json:"sentiment"`
	ID        []uint8 `json:"id"`
	Created   string  `json:"created"`
}

func OpenDBIfClosed() *sql.DB {
	var err error

	if db == nil {
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		checkErr(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS hashtags (id serial primary key, hashtag varchar(180) unique, sentiment char(9), created timestamp with time zone default current_timestamp)")
	checkErr(err)

	err = db.Ping()
	if err != nil {
		db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		checkErr(err)
	}

	return db
}

func ViewRows() []Hashtag {
	db = OpenDBIfClosed()

	rows, err := db.Query("SELECT * FROM hashtags")
	checkErr(err)
	defer rows.Close()

	var hashtags []Hashtag

	for rows.Next() {
		var hashtag Hashtag
		err = rows.Scan(&hashtag.ID, &hashtag.Name, &hashtag.Sentiment, &hashtag.Created)
		checkErr(err)
		hashtags = append(hashtags, hashtag)
	}

	return hashtags
}

func IsInTable(trend anaconda.Trend) bool {
	db = OpenDBIfClosed()

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

func InsertHashtag(hashtag string, sentiment string) {
	db = OpenDBIfClosed()

	stmt, err := db.Prepare("INSERT INTO hashtags(hashtag, sentiment) VALUES($1, $2);")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(hashtag, sentiment)
	checkErr(err)
}

func SelectRandomHashtag(sentiment string) Hashtag {
	db = OpenDBIfClosed()

	var hashtag Hashtag
	err := db.QueryRow("SELECT * FROM hashtags WHERE sentiment = $1 ORDER BY random()", sentiment).Scan(&hashtag.ID, &hashtag.Name, &hashtag.Sentiment, &hashtag.Created)
	checkErr(err)

	return hashtag
}

func PurgeDB(sentiment string) {
	db = OpenDBIfClosed()

	rows, err := db.Query("SELECT created FROM hashtags WHERE sentiment = $1 ORDER BY created DESC;", sentiment)
	checkErr(err)
	defer rows.Close()

	// get total # of rows and save created date for
	var count int
	var tenthCreatedAt time.Time
	for rows.Next() {
		count++
		if count == 10 {
			rows.Scan(&tenthCreatedAt)
		}
	}

	if count < 10 {
		return
	}

	_, err = db.Exec("DELETE FROM hashtags WHERE sentiment = $1 AND created < $2", sentiment, tenthCreatedAt)
	checkErr(err)

}
