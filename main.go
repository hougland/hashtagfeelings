package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	SetEnvVars() // from local, untracked env.go file which sets secrets

	GetTweets()
}

func checkErr(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No results found")
		} else {
			panic(err)
		}
	}
}

func openDBConnection() {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "BluePenguin", "hashtagfeelings")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
}
