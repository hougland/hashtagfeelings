package main

import "testing"

func TestGetTrends(t *testing.T) {
	SetEnvVars()

	// trends := GetTrends()
	//
	// if len(trends) != 50 {
	// 	t.Error("Expected 50, got ", len(trends))
	// }
}

func TestGetTweets(t *testing.T) {
	SetEnvVars()
	trends := GetTrends()
	tweets := GetTweets(trends[0])
	// fmt.Printf("type: %T", tweets)

	if len(tweets) > 50 {
		t.Error("Expected less than 50, got ", len(tweets))
	}
}

func TestCleanTweets(t *testing.T) {

}
