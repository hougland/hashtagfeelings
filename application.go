package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", ViewAllRows)
	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)
	http.HandleFunc("/updatehashtags", Updated)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
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
