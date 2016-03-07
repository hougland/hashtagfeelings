package main

import (
	"encoding/json"
	"net/http"
)

func Positive(w http.ResponseWriter, r *http.Request) {
	db := OpenDBConnection()
	hashtag := SelectRandomHashtag(db, "positive")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtag); err != nil {
		panic(err)
	}
}

func Negative(w http.ResponseWriter, r *http.Request) {
	db := OpenDBConnection()
	hashtag := SelectRandomHashtag(db, "negative")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtag); err != nil {
		panic(err)
	}
}

func Updated(w http.ResponseWriter, r *http.Request) {
	updateHashtags()
	db := OpenDBConnection()
	hashtags := ViewRows(db)
	defer db.Close()

	js, err := json.Marshal(hashtags)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
