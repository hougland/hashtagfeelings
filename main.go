package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "BluePenguin", "hashtagfeelings")
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	// temporarily hard coded
	// InsertHashtag(db, "#capstone")

	ShowAllHashtags(db)

	SelectOneHashtag(db)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
