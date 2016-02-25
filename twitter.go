package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
