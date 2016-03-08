package main

import (
	"encoding/json"
	"net/http"
)

func ViewAllRows(w http.ResponseWriter, r *http.Request) {
	hashtags := ViewRows()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtags); err != nil {
		panic(err)
	}
}

func Positive(w http.ResponseWriter, r *http.Request) {
	hashtag := SelectRandomHashtag("positive")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtag); err != nil {
		panic(err)
	}
}

func Negative(w http.ResponseWriter, r *http.Request) {
	hashtag := SelectRandomHashtag("negative")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hashtag); err != nil {
		panic(err)
	}
}

func Updated(w http.ResponseWriter, r *http.Request) {
	updateHashtags()

	hashtags := ViewRows()

	js, err := json.Marshal(hashtags)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func updateHashtags() {
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
