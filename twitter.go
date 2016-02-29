package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type Trend struct {
	Name            string `json:"name"`
	Query           string `json:"query"`
	Url             string `json:"url"`
	PromotedContent string `json:"promoted_content"`
}

func GetTrends() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	trends, err := api.GetTrendsByPlace(1, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", trends.Trends)
}

func GetTweets() {
	// accepts a single trends obj (?)
	// returns array (?) of popular tweets with a particular hashtag
	var trend = Trend{"#golang", "%23golang", "http://twitter.com/search?q=%23GoLang", ""}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	searchResult, err := api.GetSearch(trend.Query, nil)
	if err != nil {
		panic(err)
	}
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

}
