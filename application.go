package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", Positive)
	http.HandleFunc("/n", Negative)
	http.HandleFunc("/updatehashtags", Updated)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
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

func Positive(w http.ResponseWriter, r *http.Request) {
	db := OpenDBConnection()
	userinfo := ViewRows(db)

	js, err := json.Marshal(userinfo)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Negative(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Ricky", []string{"cats", "cats"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
