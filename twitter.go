package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

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
}
