package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	fmt.Println("starting app")
	db = OpenDBIfClosed()

	http.HandleFunc("/", ViewAllRows)
	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)
	http.HandleFunc("/updatehashtags", Updated)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			// should add an error message in the json
		} else {
			fmt.Println(err)
			panic(err)
		}
	}
}
