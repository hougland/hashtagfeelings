package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	// put the scheduling operations in their own go routines so they can occur
	// at the same time as listening/routing
	go ScheduleUpdateHashtags()
	go SchedulePurgeDB()

	db = OpenDBIfClosed()

	http.HandleFunc("/", ViewAllRows)
	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)
	http.HandleFunc("/updatehashtags", Updated)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	checkErr(err)
}

func UpdateHashtags() {
	// get trends
	trends := GetTrends()

	// for each trend, make sure it's not in db, get its tweets, run sentiment analysis, save in db
	for _, trend := range trends {
		if IsInTable(trend) == false {
			tweets := GetTweets(trend)
			isSentimental, whichSentiment := SentimentAnalysis(tweets)
			if isSentimental {
				InsertHashtag(trend.Name, whichSentiment)
			}
		}
	}
}

// TODO: print more details about errors
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
