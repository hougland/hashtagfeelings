package main

import (
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestCreateSentimentQuery(t *testing.T) {
	var tweet1, tweet2 anaconda.Tweet
	tweet1.Text = "test tweet"
	tweet2.Text = ""

	var tweetSlice1, tweetSlice2 []anaconda.Tweet
	tweetSlice1 = append(tweetSlice1, tweet1)
	tweetSlice2 = append(tweetSlice2, tweet2)

	var query1, query2 SentimentQuery
	query1.Data = append(query1.Data, FormatTweet(tweet1))
	query2.Data = append(query2.Data, FormatTweet(tweet2))

	cases := []struct {
		in   []anaconda.Tweet
		want SentimentQuery
	}{
		{tweetSlice1, query1},
		{tweetSlice2, query2},
	}

	for _, c := range cases {
		got := CreateSentimentQuery(c.in)
		if got.Data[0] != c.want.Data[0] {
			t.Errorf("CreateSentimentQuery(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

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
			t.Errorf("FormatTweet(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

// how do I not send to the actual API?
func TestSentimentAnalysis(t *testing.T) {
	var tweet1, tweet2, tweet3, tweet4 anaconda.Tweet
	tweet1.Text = "happy happy happy"
	tweet2.Text = "sad sad sad"
	tweet3.Text = "neutral neutral neutral"
	tweet4.Text = ""

	var tweetSlice1, tweetSlice2, tweetSlice3, tweetSlice4 []anaconda.Tweet
	tweetSlice1 = append(tweetSlice1, tweet1)
	tweetSlice2 = append(tweetSlice2, tweet2)
	tweetSlice3 = append(tweetSlice3, tweet3)
	tweetSlice4 = append(tweetSlice4, tweet4)

	cases := []struct {
		in    []anaconda.Tweet
		want1 bool
		want2 string
	}{
		{tweetSlice1, true, "positive"},
		{tweetSlice2, true, "negative"},
		{tweetSlice3, false, ""},
		{tweetSlice4, false, ""},
	}

	for _, c := range cases {
		got1, got2 := SentimentAnalysis(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("SentimentAnalysis(%v) == %v and %v, want1 %v, want2 %v", c.in[0].Text, got1, got2, c.want1, c.want2)
		}
	}
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
			t.Errorf("IsSentimental(%v) == %v and %v, want1 %v, want2 %v", c.in, got1, got2, c.want1, c.want2)
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
			t.Errorf("GetScore(%v) == %v, want %v", c, score, c.want)
		}
	}
}
