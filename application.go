package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	SetEnvVars() // from local, untracked env.go file which sets secrets

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}

func updateHashtags() {
	// open db
	db := OpenDBConnection()
	defer db.Close()

	// get trends
	trends := GetTrends()

	// for each trend, make sure it's not in db, get its tweets, run sentiment analysis, save in db
	for _, trend := range trends {
		if IsInTable(db, trend) == false {
			tweets := GetTweets(trend)
			isSentimental, whichSentiment := SentimentAnalysis(tweets)
			if isSentimental {
				InsertHashtag(db, trend.Name, whichSentiment)
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
