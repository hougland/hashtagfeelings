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
	Data [][]byte `json:"data"`
}

func SentimentAnalysis(tweet anaconda.Tweet) {
	// returns sentiment object (?) - positive, negative, and an intensity of sentiment

	jsonStr := FormatTweet(tweet)

	url := "http://www.sentiment140.com/api/bulkClassifyJson"
	fmt.Println("URL:>", url)

	// var jsonStr = []byte(`{"data":[{"text":"Buy cheese and bread for breakfast."}]}`)
	fmt.Println(string(jsonStr))
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

func FormatTweet(tweet anaconda.Tweet) []byte {
	tweetStruct := &TweetText{Text: tweet.Text}

	marshaledTweet, err := json.Marshal(tweetStruct)
	if err != nil {
		panic(err)
	}

	return marshaledTweet
}

// func BuildSentimentQuery(tweetSlice []anaconda.Tweet) [][]byte {
// 	var querySlice [][]byte
//
// 	for _, tweet := range tweetSlice {
// 		marshaledTweet := FormatTweet(tweet)
// 		querySlice = append(querySlice, marshaledTweet)
// 	}
//
// 	for _, query := range querySlice {
// 		fmt.Println(string(query))
// 	}
//
// 	return querySlice
// }
