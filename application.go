package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// SetEnvVars() // from local, untracked env.go file which sets secrets
	// updateHashtags()

	f, _ := os.Create("/var/log/golang/golang-server.log")
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/positive", Positive)
	http.HandleFunc("/negative", Negative)

	log.Printf("Listening on port 5000")
	http.ListenAndServe(":5000", nil)

}

// func updateHashtags() {
// 	// open db
// 	db := OpenDBConnection()
// 	defer db.Close()
//
// 	// get trends
// 	trends := GetTrends()
//
// 	// for each trend, make sure it's not in db, get its tweets, run sentiment analysis, save in db
// 	for _, trend := range trends {
// 		if IsInTable(db, trend) == false {
// 			tweets := GetTweets(trend)
// 			isSentimental, whichSentiment := SentimentAnalysis(tweets)
// 			if isSentimental {
// 				InsertHashtag(db, trend.Name, whichSentiment)
// 			}
// 		}
// 	}
// }

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
