package main

import (
	"testing"
)

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

}
