package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func GetTrends() []anaconda.Trend {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	trendResponse, err := api.GetTrendsByPlace(1, nil)
	if err != nil {
		panic(err)
	}

	return trendResponse.Trends
}

func GetTweets(trend anaconda.Trend) []anaconda.Tweet {
	// accepts a single trends obj
	// returns array of popular tweets with a particular hashtag

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

	return searchResult.Statuses

}

// func CleanTweets(SearchResponse) {
// 	// accepts a single
// }
