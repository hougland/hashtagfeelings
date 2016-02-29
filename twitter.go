package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
)

func GetWorldWideTrends() {
	// not sure why Getenv doesn't work - need to investigate
	// consumerKey := os.Getenv("CONSUMER_KEY")
	// consumerSecret := os.Getenv("CONSUMER_SECRET")
	// accessToken := os.Getenv("ACCESS_TOKEN")
	// accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	// this works when hard coded. YAY!
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")
	trends, err := api.GetTrendsByPlace(1, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", trends.Trends)
}

func GetTrends() {
	// returns array of worldwide trends
	response, err := http.Get("https://api.twitter.com/1.1/trends/place.json?id=1")
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			panic(err)
		}
		fmt.Printf("%s\n", string(contents))
	}

}

func GetTweets() {
	// accepts a single trends obj (?)
	// returns array (?) of popular tweets with a particular hashtag
}
