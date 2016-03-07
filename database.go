package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
)

type Hashtag struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
	ID        int    `json:"id"`
}

type Userinfo struct {
	Uid        int       `json:"uid"`
	Username   string    `json:"username"`
	Department string    `json:"department"`
	Created    time.Time `json:"created"`
}

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)

	return db
}

func ViewRows(db *sql.DB) []Userinfo {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	defer rows.Close()

	var userinfos []Userinfo

	for rows.Next() {
		var userinfo Userinfo
		err = rows.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Department, &userinfo.Created)
		checkErr(err)
		userinfos = append(userinfos, userinfo)
	}

	return userinfos
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
