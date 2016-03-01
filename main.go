package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	SetEnvVars() // from local, untracked env.go file which sets secrets

	trends := GetTrends()
	fmt.Println(len(trends))
	fmt.Println(trends[0])
	tweets := GetTweets(trends[0])
	fmt.Println(len(tweets))
	fmt.Println(tweets[0])

	// tweets := GetTweets(trends[0])
	// fmt.Println(trends[0].Name)
	// fmt.Println(len(tweets))
	//
	// for _, tweet := range tweets {
	// 	fmt.Println(tweet.Text)
	// }

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

func updateHashtags() {
	// open db
	// db := OpenDBConnection()

	// get trends
	trends := GetTrends()

	// get tweets for each trend
	for _, trend := range trends {
		GetTweets(trend)
	}

	// clean tweets for each trend

	// send tweets for each trend to sentiment analysis api

	// check if sentiment is strong enough to save in db

	// save in db

}
