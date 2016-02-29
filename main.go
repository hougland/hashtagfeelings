package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	SetEnvVars() // from local, untracked env.go file which sets secrets

	// Open connection to db:
	// dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "BluePenguin", "hashtagfeelings")
	// db, err := sql.Open("postgres", dbinfo)
	// checkErr(err)
	// defer db.Close()

	fmt.Println("Testing:", os.Getenv("CONSUMER_KEY"))

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
