package main

import (
	"net/url"
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
	// currently returns 15 tweets - need to make them the popular ones

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}
	v.Set("result_type", "popular")
	v.Set("lang", "en")
	v.Set("count", "50")

	searchResult, err := api.GetSearch(trend.Query, v)
	if err != nil {
		panic(err)
	}

	return searchResult.Statuses
}

// func CleanTweets(tweets []anaconda.Tweet) []anaconda.Tweet {
// acceps slice of Tweets
// returns a slice of "clean" Tweets (remove special characters, etc.)
// Tweets ready to send to sentiment analysis
// }
