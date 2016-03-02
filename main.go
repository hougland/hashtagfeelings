package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	updateHashtags()
}

func updateHashtags() {
	// open db
	db := OpenDBConnection()
	defer db.Close()

	// get trends
	trends := GetTrends()

	// for each trend, get its tweets, run sentiment analysis, save in db
	for _, trend := range trends {
		tweets := GetTweets(trend)
		isSentimental, whichSentiment := SentimentAnalysis(tweets)
		if isSentimental {
			InsertHashtag(db, trend.Name, whichSentiment)
		}
	}
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
