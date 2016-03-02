package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
)

type TweetText struct {
	Text string `json:"text"`
}

type SentimentQuery struct {
	Data []*TweetText `json:"data"`
}

func CreateSentimentQuery(tweets []anaconda.Tweet) SentimentQuery {
	var query SentimentQuery

	for _, tweet := range tweets {
		query.Data = append(query.Data, FormatTweet(tweet))
	}

	return query
}

func FormatTweet(tweet anaconda.Tweet) *TweetText {
	tweetStruct := &TweetText{Text: tweet.Text}

	return tweetStruct
}

func SentimentAnalysis(tweets []anaconda.Tweet) {
	// returns sentiment object (?) - positive, negative, and an intensity of sentiment
	query := CreateSentimentQuery(tweets)
	jsonStr, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	url := "http://www.sentiment140.com/api/bulkClassifyJson"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

// func IsSentimental() {
// 	// accepts sentiment object, returns true/false based on if semtiment strong enough to save
// }
