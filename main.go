package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	SetEnvVars() // from local, untracked env.go file which sets secrets

	db := OpenDBConnection()
	defer db.Close()

}

func updateHashtags() {
	// open db
	db := OpenDBConnection()
	defer db.Close()

	// get trends
	trends := GetTrends()

	// for each trend, get it's tweets, run sentiment analysis, save in db
	for _, trend := range trends {
		tweets := GetTweets(trend)
		isSentimental, whichSentiment := SentimentAnalysis(tweets)
		if isSentimental && whichSentiment == "positive" {
			// save in pos_hashtags table
		} else if isSentimental && whichSentiment == "negative" {
			// save in neg_hashtags table
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
