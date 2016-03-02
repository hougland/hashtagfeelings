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
	Text     string        `json:"text"`
	Polarity int           `json:"polarity"`
	Meta     MetaSentiment `json:"meta"`
}

type MetaSentiment struct {
	Language string `json:"language"`
}

type SentimentQuery struct {
	Data []TweetText `json:"data"`
}

func CreateSentimentQuery(tweets []anaconda.Tweet) SentimentQuery {
	var query SentimentQuery

	for _, tweet := range tweets {
		query.Data = append(query.Data, FormatTweet(tweet))
	}

	return query
}

func FormatTweet(tweet anaconda.Tweet) TweetText {
	tweetStruct := TweetText{Text: tweet.Text}

	return tweetStruct
}

func SentimentAnalysis(tweets []anaconda.Tweet) SentimentQuery {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var unmarshaledQuery SentimentQuery
	err = json.Unmarshal(body, &unmarshaledQuery)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return unmarshaledQuery
}

func IsSentimental(total float64) (bool, string) {
	// accepts sentiment object, returns true/false based on if semtiment strong enough to save
	if total >= 3 {
		return true, "positive"
	} else if total <= 1 {
		return true, "negative"
	} else {
		return false, ""
	}
}

func GetScore(sentimentObj SentimentQuery) float64 {
	var (
		numTweets float64
		scores    float64
	)

	numTweets = float64(len(sentimentObj.Data))

	for _, tweet := range sentimentObj.Data {
		scores += float64(tweet.Polarity)
	}

	total := scores / numTweets
	fmt.Printf("numTweets: %v", numTweets)
	fmt.Printf("scores: %v", scores)
	fmt.Printf("total: %v", total)

	return total
}
