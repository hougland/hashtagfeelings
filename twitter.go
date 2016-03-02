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
	defer api.Close()

	trendResponse, err := api.GetTrendsByPlace(23424977, nil)
	if err != nil {
		panic(err)
	}

	return trendResponse.Trends
}

func GetTweets(trend anaconda.Trend) []anaconda.Tweet {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CUSTOMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	defer api.Close()

	v := url.Values{}
	v.Set("result_type", "mixed")
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

// example of testing values of a type
// func (signature *Signature) valid() bool {
//     return len(signature.FirstName) > 0 &&
//         len(signature.LastName) > 0 &&
//         len(signature.Email) > 0 &&
//         signature.Age >= 18 && signature.Age <= 180 &&
//         len(signature.Message) < 140
// }
// }

// func CleanTrends() {
// 	// accepts single trend, verifies that it's english
// }
