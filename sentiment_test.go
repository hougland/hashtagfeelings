package main

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestCreateSentimentQuery(t *testing.T) {
	// var tweet1 anaconda.Tweet
	// tweet1.
	// var tweetSlice []anaconda.Tweet
}

// func CreateSentimentQuery(tweets []anaconda.Tweet) SentimentQuery {
// 	var query SentimentQuery
//
// 	for _, tweet := range tweets {
// 		query.Data = append(query.Data, FormatTweet(tweet))
// 	}
//
// 	return query
// }

func TestFormatTweet(t *testing.T) {
	var tweet1, tweet2 anaconda.Tweet
	tweet1.Text = "test tweet"
	tweet2.Text = ""

	var tweetStruct1, tweetStruct2 TweetText
	tweetStruct1.Text = "test tweet"
	tweetStruct2.Text = ""

	cases := []struct {
		in   anaconda.Tweet
		want TweetText
	}{
		{tweet1, tweetStruct1},
		{tweet2, tweetStruct2},
	}

	for _, c := range cases {
		got := FormatTweet(c.in)
		if got != c.want {
			t.Errorf("FormatTweet(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// func FormatTweet(tweet anaconda.Tweet) TweetText {
// 	tweetStruct := TweetText{Text: tweet.Text}
//
// 	return tweetStruct
// }

func TestSentimentAnalysis(t *testing.T) {

}

func TestIsSentimental(t *testing.T) {
	cases := []struct {
		in    float64
		want1 bool
		want2 string
	}{
		{4, true, "positive"},   // very positive
		{0, true, "negative"},   // very negative
		{2, false, ""},          // very neutral
		{2.9, true, "positive"}, // barely positive
		{1.4, true, "negative"}, // barely negative
		{2.8, false, ""},        // barely neutral
		{1.5, false, ""},        // barely neutral
	}

	for _, c := range cases {
		got1, got2 := IsSentimental(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("IsSentimental(%q) == %q and %q, want1 %q, want2 %q", c.in, got1, got2, c.want1, c.want2)
		}
	}
}

func TestGetScore(t *testing.T) {
	var tweet1, tweet2, tweet3 TweetText
	tweet1.Text = "happy happy happy"
	tweet1.Polarity = 4

	tweet2.Text = "sad sad sad"
	tweet2.Polarity = 0

	tweet3.Text = "neutral neutral neutral"
	tweet3.Polarity = 2

	var sentimentObj1, sentimentObj2, sentimentObj3, sentimentObj4 SentimentQuery
	sentimentObj1.Data = append(sentimentObj1.Data, tweet1)
	sentimentObj2.Data = append(sentimentObj2.Data, tweet2)
	sentimentObj3.Data = append(sentimentObj3.Data, tweet3)
	sentimentObj4.Data = append(sentimentObj4.Data, tweet1)
	sentimentObj4.Data = append(sentimentObj4.Data, tweet2)

	cases := []struct {
		in   SentimentQuery
		want float64
	}{
		{sentimentObj1, 4},
		{sentimentObj2, 0},
		{sentimentObj3, 2},
		{sentimentObj4, 2},
	}

	for _, c := range cases {
		score := GetScore(c.in)
		if score != c.want {
			t.Errorf("GetScore(%q) == %q, want %q", c, score, c.want)
		}
	}
}
