package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func connectWithApi() *anaconda.TwitterApi {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	return api
}

func GetTrends() []anaconda.Trend {
	api := connectWithApi()
	defer api.Close()

	trendResponse, err := api.GetTrendsByPlace(23424977, nil)
	if err != nil {
		panic(err)
	}

	return trendResponse.Trends
}

func GetTweets(trend anaconda.Trend) []anaconda.Tweet {
	api := connectWithApi()
	defer api.Close()

	v := url.Values{}
	v.Set("result_type", "mixed")
	v.Set("lang", "en")
	v.Set("count", "100")

	searchResult, err := api.GetSearch(trend.Query, v)
	if err != nil {
		panic(err)
	}

	return searchResult.Statuses
}
