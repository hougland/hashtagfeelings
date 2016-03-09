package main

import "testing"

func TestCreateSentimentQuery(t *testing.T) {

}

func TestFormatTweet(t *testing.T) {

}

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
